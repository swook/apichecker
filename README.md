API Checker
===========

API Checker is a Google App Engine application for checking API compatibility/completeness of applications in development, written in Go.

API profiles (`APIProfile`) can be created with comparison cases (`Comp`) using either of JSON, XPath,
and Header matching methods.

Application profiles (`Application`), when referred to an `APIProfile` can be used to check and monitor the API completeness of the application.

Documentation is available at [http://godoc.org/github.com/swook/apichecker](http://godoc.org/github.com/swook/apichecker)

The application runs at  [http://api-checker.appspot.com/](http://api-checker.appspot.com/)

This application is still in development.
