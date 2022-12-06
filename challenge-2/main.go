package main

import (
	"fmt"
	"github.com/syke99/waggy"
	"net/http"
	"strings"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
)

var flg waggy.FullCGI

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		/*
			Create a WaggyHandler to handle the hello endpoint
		*/
		elfName := waggy.Vars(r)["name"]

		// This will need to be done by the outbound_http requirement for Challenge 2
		elfName = strings.ToLower(elfName)

		helloWorldElfBytes := []byte(fmt.Sprintf("{ \"message\":\"Hello, %s\" }", elfName))

		helloHandler := func(w http.ResponseWriter, r *http.Request) {
			// call the lowering endpoint here
			w.Header().Set("Content-Type", "application/json")
			w.Write(helloWorldElfBytes)
		}

		helloWaggyHandler := waggy.InitHandler(nil).
			MethodHandler(http.MethodGet, helloHandler)

		/*
			Create a WaggyHandler to handle the lowering endpoint
		*/
		lowerCaseHandler := func(w http.ResponseWriter, r *http.Request) {
			// This is where lowering the string will take place
		}

		lowerCaseWaggyHandler := waggy.InitHandler(nil).
			MethodHandler(http.MethodPost, lowerCaseHandler)

		/*
			Create a WaggyHandler to handle the error endpoint
		*/
		errorHandler := func(w http.ResponseWriter, r *http.Request) {
			waggy.WriteDefaultResponse(w, r)
		}

		wError := waggy.WaggyError{
			Type:     "/",
			Title:    "Error route on root",
			Detail:   "route not found",
			Status:   0,
			Instance: "/",
		}

		// Make sure all HTTP methods return the errorHandler whenever this
		// WaggyHandler is served
		errorWaggyHandler := waggy.InitHandler(nil).
			WithDefaultErrorResponse(wError, http.StatusNotFound).
			MethodHandler(http.MethodGet, errorHandler).
			MethodHandler(http.MethodPost, errorHandler).
			MethodHandler(http.MethodPut, errorHandler).
			MethodHandler(http.MethodDelete, errorHandler).
			MethodHandler(http.MethodConnect, errorHandler).
			MethodHandler(http.MethodHead, errorHandler).
			MethodHandler(http.MethodOptions, errorHandler).
			MethodHandler(http.MethodPatch, errorHandler).
			MethodHandler(http.MethodTrace, errorHandler)

		/*
			Create a WaggyRouter to handle all of the routes
			with their appropriate WaggyHandlers
		*/
		router := waggy.InitRouter(&flg).
			Handle("/", errorWaggyHandler).
			Handle("/hello/{name}", helloWaggyHandler).
			Handle("/lowercase", lowerCaseWaggyHandler)

		router.ServeHTTP(w, r)
	})
}

func main() {}
