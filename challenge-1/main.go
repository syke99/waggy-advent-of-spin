package main

import (
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/syke99/waggy"
)

var flg waggy.FullCGI

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		/*
			Create a WaggyHandler to handle the hello endpoint
		*/
		helloWorldJSONBytes := []byte("{ \"message\": \"Hello, world!\" }")

		helloHandler := func(w http.ResponseWriter, r *http.Request) {
			waggy.WriteDefaultResponse(w, r)
		}

		handler := waggy.InitHandlerWithRoute("/hello", &flg).
			WithDefaultResponse("application/json", helloWorldJSONBytes).
			MethodHandler(http.MethodGet, helloHandler)

		handler.ServeHTTP(w, r)
	})
}

func main() {}
