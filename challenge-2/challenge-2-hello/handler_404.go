package main

import (
	"fmt"
	"github.com/syke99/waggy"
	"net/http"
)

func Handler404(flg waggy.FullServer) *waggy.Handler {
	_handler404 := func(w http.ResponseWriter, r *http.Request) {
		// if the request path is "/", return a 404
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "resource not found")
	}

	// initialize a handler with Waggy, with the route template that
	// should be used with the request (putting path variables inside
	// of curly brackets like {name} allows you to use waggy.Vars(r)
	// to return a map of any matched path variables; so calling
	// waggy.Vars(r)["name"] will return the section of the path
	// that matches {name}, if any) and map the handler func to the
	// appropriate HTTP Method (in more complex examples, this allows
	// one handler initialized with Waggy to handle different HTTP Methods
	// with different handler funcs
	handler := waggy.InitHandlerWithRoute("/", &flg).
		// using waggy.AllHTTPMethods makes it clean and simple to map
		// a single handler to all HTTP methods, such as this 404 handler
		MethodHandler(waggy.AllHTTPMethods(), _handler404)

	return handler
}
