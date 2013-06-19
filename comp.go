// Copyright 2013 The apichecker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package apichecker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"launchpad.net/xmlpath"
	"net/http"
	"regexp"
	"strings"
)

const (
	CompJSON = iota
	CompXPath
	CompHeader
)

// Profile defines the expected responses of an API
type APIProfile struct {
	ShortName string
	Hosts     []string
	Comps     []Comp
}

// Application defines an application which implements an API
type Application struct {
	ShortName string
	API       string // APIProfile.ShortName
	Host      string
}

// Comp is a struct which defines a comparison case
//
type Comp struct {
	QueryPath   string
	Type        int
	Description string
	Path        string
	DataMatch   string
}

func (c *Comp) dataMatch(data string) bool {
	ok, err := regexp.MatchString(c.DataMatch, data)
	if err != nil {
		fmt.Printf("dataMatch (%v) & (%v): %v\n", c.DataMatch, data, err)
		return false
	}
	return ok
}

func (c *Comp) doCompJSON(r *http.Response) bool {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("doCompJSON (ioutil.ReadAll): %v\n", err)
		return false
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("doCompJSON (json.Unmarshal): %v\n", err)
		return false
	}

	// Convert paths to []string
	// paths are stored as a newline (\n) delimited key tree
	nodes := strings.Split(c.Path, "\n")
	for i, v := range nodes {
		nodes[i] = strings.Trim(v, " \t")
	}

	ok := c.traverseMatchJSON(nodes, data)
	return ok
}

func (c *Comp) traverseMatchJSON(nodes []string, data interface{}) bool {
	switch data.(type) {
	case bool:
		d := fmt.Sprintf("%v", data.(bool))
		return c.dataMatch(d)
	case float64:
		d := fmt.Sprintf("%v", data.(float64))
		return c.dataMatch(d)
	case string:
		d := fmt.Sprintf("%v", data.(string))
		return c.dataMatch(d)
	case map[string]interface{}:
		if len(nodes) == 0 {
			return false
		}
		d := data.(map[string]interface{})
		return c.traverseMatchJSON(nodes[1:], d[nodes[0]])
	}
	return false
}

func (c *Comp) doCompXPath(r *http.Response) bool {
	path := xmlpath.MustCompile(c.Path)
	root, err := xmlpath.ParseHTML(r.Body)
	if err != nil {
		fmt.Printf("doCompXPath: %v\n", err)
		return false
	}

	value, ok := path.String(root)
	if !ok {
		return false
	}

	ok = c.dataMatch(value)
	return ok
}

func (c *Comp) doCompHeader(r *http.Response) bool {
	value := r.Header.Get(c.Path)

	ok := c.dataMatch(value)
	return ok
}
