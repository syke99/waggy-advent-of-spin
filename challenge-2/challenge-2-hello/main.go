package main

import (
	"bytes"
	"fmt"
	"github.com/syke99/waggy"
	"io/ioutil"
	"net/http"
	"strings"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
)

var flg waggy.FullCGI

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		name := waggy.Vars(r)["name"]

		helloElfJSONBytes := []byte("{ \"message\":\"Hello, world\" }")

		if name != "" {
			lowerCaseName := ""

			resp, err := spinhttp.Post("127.0.0.1/lowercase", "application/json", bytes.NewBufferString(fmt.Sprintf("{ \"message\": %s", name)))
			if err != nil {
				waggy.WriteDefaultErrorResponse(w, r)
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				waggy.WriteDefaultErrorResponse(w, r)
			}

			lowerCaseName = strings.Split(string(body), " ")[2:3][0]

			if lowerCaseName[:1] == "\"" &&
				lowerCaseName[len(lowerCaseName)-1:] == "\"" {
				lowerCaseName = lowerCaseName[1 : len(lowerCaseName)-1]
			}

			helloElfJSONBytes = []byte(fmt.Sprintf("{ \"message\":\"Hello, %s\" }", lowerCaseName))
		}

		helloHandler := func(w http.ResponseWriter, r *http.Request) {
			waggy.WriteDefaultResponse(w, r)
		}

		we := waggy.WaggyError{
			Title:    "Internal server error",
			Detail:   "sorry, something went wrong!",
			Status:   0,
			Instance: "/",
		}

		handler := waggy.InitHandlerWithRoute("/hello", &flg).
			WithDefaultResponse(helloElfJSONBytes).
			WithDefaultErrorResponse(we, http.StatusInternalServerError).
			MethodHandler(http.MethodGet, helloHandler)

		handler.ServeHTTP(w, r)
	})
}

func main() {}
