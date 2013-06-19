// Copyright 2013 The apichecker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package apichecker is an App Engine application for checking
// API compatibility or completeness of applications in development
package apichecker

import (
	// "appengine"
	// "appengine/datastore"
	"html/template"
	"net/http"
)

const (
	templateDir = "template/"
)

// Parse required templates
var (
	mainTemplate     = template.Must(template.ParseFiles(templateDir + "main.html"))
	notFoundTemplate = template.Must(template.ParseFiles(templateDir + "404.html"))
	contactTemplate  = template.Must(template.ParseFiles(templateDir + "contact.html"))
)

// Initialise all handlers
func init() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/get", mainHandler)
	http.HandleFunc("/contact", contactHandler)
}

// mainHandler deals with all requests for path "/"
func mainHandler(w http.ResponseWriter, r *http.Request) {
	// c := appengine.NewContext(r)

	if r.URL.Path != "/" {
		notFoundHandler(w, r)
		return
	}

	if err := mainTemplate.Execute(w, r.URL); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// notFoundHandler deals with 404 errors
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	if err := notFoundTemplate.Execute(w, r.URL); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// getHandler responds to GET requests to /get/{Application.ShortName}
func getHandler(w http.ResponseWriter, r *http.Request) {

}

// contactHandler deals with all requests for path "/about"
func contactHandler(w http.ResponseWriter, r *http.Request) {
	if err := contactTemplate.Execute(w, r.URL); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
