// Copyright 2013 The apichecker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package apichecker

// mlabns_py is an API profile for the Python version of mlab-ns
// developed by Klaudiu Perta as part of GSoC 2012
var mlabns_py = APIProfile{
	ShortName: "mlab-ns",
	Comps: []Comp{
		// Test re-direct
		Comp{
			QueryPath:   "",
			Type:        CompHeader,
			Description: "",
			Path:        "",
			DataMatch:   "",
		},
	},
}

// mlabns_appspot is the Python version of mlab-ns developed by
// Klaudiu Perta as part of GSoC 2012
var mlabns_appspot = Application{
	ShortName: "mlab-ns-py",
	API:       "mlab-ns",
	Host:      "http://mlab-ns.appspot.com/",
}
