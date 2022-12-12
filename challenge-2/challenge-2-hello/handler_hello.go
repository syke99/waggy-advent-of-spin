package main

import (
	"bytes"
	"fmt"
	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/syke99/waggy"
	"io/ioutil"
	"net/http"
	"strings"
)

func HelloHandler(flg waggy.FullServer) *waggy.Handler {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		// use Waggy to attempt to grab the provided name
		// to say hello to
		name := waggy.Vars(r)["name"]

		// start off assuming that no name was provided
		helloElfJSON := "{ \"message\": \"Hello, world\" }"

		// if there was a name provided, make an outbound http request to the lower casing service
		// to make sure the name is all lowercase
		if name != "" {
			lowerCaseName := ""

			// build the JSON request with the name from the incoming request to make a request
			// to the lower casing service
			reqBody := bytes.NewBufferString(fmt.Sprintf("{ \"message\": \"%s\" }", name))

			// call the lower casing service with the built request body
			resp, err := spinhttp.Post("http://127.0.0.1:3001/lowercase", "application/json", reqBody)
			if err != nil {
				// if the request errors, write the default response that was set with the handler
				// with Waggy
				waggy.WriteDefaultErrorResponse(w, r)
				return
			}
			defer resp.Body.Close()

			// make a byte slice to read the resp.Body into
			body := make([]byte, 0)

			if resp != nil {
				// read the resp.Body
				body, err = ioutil.ReadAll(resp.Body)
			}

			if err != nil {
				// if reading the resp.Body errors, write the default response that was set with
				// the handler with Waggy
				waggy.WriteDefaultErrorResponse(w, r)
				return
			}

			// extract the lower cased name from the read resp.Body
			splitBody := strings.Split(string(body), ":")

			if len(splitBody) == 2 {
				// after splitting the response body, remove the trailing "}"
				// from the second value in splitBody and set it equal to lowerCaseName
				lowerCaseName = strings.Replace(splitBody[1], "}", "", -1)
				// strip any quotation marks from lowerCaseName
				lowerCaseName = strings.Replace(lowerCaseName, "\"", "", -1)
				// remove any remaining whitespace from the lowerCaseName
				lowerCaseName = strings.TrimSpace(lowerCaseName)
			}

			// update the response to be written back to the client with the lower
			// cased name they provided
			helloElfJSON = fmt.Sprintf("{ \"message\": \"Hello, %s\" }", lowerCaseName)
		}
		// set the Content-Type header and write the response back to the client
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, helloElfJSON)
	}

	// set up a WaggyError to be used with waggy.WriteDefaultErrorResponse(w, r)
	we := waggy.WaggyError{
		Title:    "Internal server error",
		Detail:   "sorry, something went wrong!",
		Status:   0,
		Instance: "/",
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
	handler := waggy.InitHandler(&flg).
		WithDefaultErrorResponse(we, http.StatusInternalServerError).
		MethodHandler(http.MethodGet, helloHandler)

	return handler
}
