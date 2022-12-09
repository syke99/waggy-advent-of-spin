package main

import (
	"github.com/syke99/waggy"
	"net/http"
)

func HelloHandler(flg waggy.FullCGI) *waggy.WaggyHandler {
	helloWorldJSONBytes := []byte("{ \"message\": \"Hello, world!\" }")

	// create a handler func that you can map to the HTTP Methods
	// you want this handler func to run on (in more complex examples, this allows
	// one handler initialized with Waggy to handle different HTTP Methods
	// with different handler funcs)
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		waggy.WriteDefaultResponse(w, r)
	}

	// initialize a handler with Waggy, with the route template that
	// should be used with the request, and map the handler func to the
	// appropriate HTTP Method (in more complex examples, this allows
	// one handler initialized with Waggy to handle different HTTP Methods
	// with different handler funcs)
	handler := waggy.InitHandlerWithRoute("/hello", &flg).
		WithDefaultResponse("application/json", helloWorldJSONBytes).
		MethodHandler(http.MethodGet, helloHandler)

	return handler
}
