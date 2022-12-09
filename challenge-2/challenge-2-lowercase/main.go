package main

import (
	"fmt"
	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/syke99/waggy"
	"io/ioutil"
	"net/http"
	"strings"
)

var flg waggy.FullCGI

func LowerCaseHandler(flg waggy.FullCGI) *waggy.WaggyHandler {
	// create a handler func that you can map to the HTTP Methods
	// you want this handler func to run on (in more complex examples, this allows
	// one handler initialized with Waggy to handle different HTTP Methods
	// with different handler funcs)
	lowerCaserHandler := func(w http.ResponseWriter, r *http.Request) {
		// read the incoming request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// if reading r.Body errors, write the default response that was set with
			// the handler with Waggy
			waggy.WriteDefaultErrorResponse(w, r)
			return
		}

		bodyString := string(body)

		// extract the name from the incoming body (this example does not shield against
		// unexpected JSON)
		bodyParts := strings.Split(bodyString, ": ")

		// make sure the name is in lower case
		lowerCasedName := strings.TrimSpace(strings.ToLower(bodyParts[1]))

		// set the Content-Type header and write the response back to the client
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fmt.Sprintf("{ \"message\": \"%s\" }", lowerCasedName))
	}

	// set up a WaggyError to be used with waggy.WriteDefaultErrorResponse(w, r)
	we := waggy.WaggyError{
		Title:    "Internal server error",
		Detail:   "sorry, something went wrong!",
		Status:   0,
		Instance: "/",
	}

	// initialize a handler with Waggy, with the route template that
	// should be used with the request, and map the handler func to the
	// appropriate HTTP Method (in more complex examples, this allows
	// one handler initialized with Waggy to handle different HTTP Methods
	// with different handler funcs)
	handler := waggy.InitHandlerWithRoute("/lowercase", &flg).
		WithDefaultErrorResponse(we, http.StatusInternalServerError).
		MethodHandler(http.MethodPost, lowerCaserHandler)

	return handler
}

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		handler := LowerCaseHandler(flg)

		// pass the incoming request into the handler initialized above
		handler.ServeHTTP(w, r)
	})
}

func main() {}
