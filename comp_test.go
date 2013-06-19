// Copyright 2013 The apichecker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package apichecker

import (
	"io"
	"net/http"
	"testing"
)

var cJSONSuccess = Comp{
	Type: CompJSON,
	Path: `A
	B
	C`,
	DataMatch: "1",
}

var cJSONFail = Comp{
	Type: CompJSON,
	Path: `A
	B
	C`,
	DataMatch: "A",
}

var dJSON = `{"A": {"B": {"C": 1, "D": 2, "E": 3}}, "F": "GHI", "G": true}`

type reader struct {
	bytes []byte
	index int
}

func (r *reader) Read(out []byte) (int, error) {
	if r.index < len(r.bytes) {
		out[0] = r.bytes[r.index]
		r.index++
		return 1, nil
	}
	r.index = 0
	return 0, io.EOF
}

func (r *reader) Close() error {
	return nil
}

func newReader(s string) *reader {
	return &reader{
		bytes: []byte(s),
		index: 0,
	}
}

func TestDoCompJSON(t *testing.T) {
	r := &http.Response{
		Body: newReader(dJSON),
	}
	ok := cJSONSuccess.doCompJSON(r)
	if !ok {
		t.Error("cJSONSuccess failed")
	}
	ok = cJSONFail.doCompJSON(r)
	if ok {
		t.Error("cJSONFail failed")
	}
}

var cXPathSuccess = Comp{
	Type:      CompXPath,
	Path:      "//ul/li[1]",
	DataMatch: "Jacob",
}

var cXPathFail = Comp{
	Type:      CompXPath,
	Path:      "//ul/li[1]",
	DataMatch: "Mary",
}

var dHTML = `<html>
	<head></head>
	<body>
		<div class="container">
			Name of students
			<ul id="students">
				<li class="name">Jacob</li>
				<li class="name">Peter</li>
				<li class="name">Mary</li>
			</ul>
		</div>
	</body>
</html>`

func TestDoCompXPath(t *testing.T) {
	r := &http.Response{Body: newReader(dHTML)}
	ok := cXPathSuccess.doCompXPath(r)
	if !ok {
		t.Error("cXPathSuccess failed")
	}
	ok = cXPathFail.doCompXPath(r)
	if ok {
		t.Error("cXPathFail failed")
	}
}

var cHeaderSuccess = Comp{
	Type:      CompHeader,
	Path:      "A",
	DataMatch: "B",
}

var cHeaderFail = Comp{
	Type:      CompHeader,
	Path:      "A",
	DataMatch: "A",
}

func TestDoCompHeader(t *testing.T) {
	h := http.Header(make(map[string][]string))
	h["A"] = []string{"B"}
	r := &http.Response{Header: h}
	ok := cHeaderSuccess.doCompHeader(r)
	if !ok {
		t.Error("cHeaderSuccess failed")
	}
	ok = cHeaderFail.doCompHeader(r)
	if ok {
		t.Error("cHeaderFail failed")
	}
}
