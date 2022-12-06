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

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		lowerCaserHandler := func(w http.ResponseWriter, r *http.Request) {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				waggy.WriteDefaultErrorResponse(w, r)
				return
			}

			bodyString := string(body)

			bodyParts := strings.Split(bodyString, ": ")

			lowerCasedName := strings.TrimSpace(strings.ToLower(bodyParts[1]))

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, fmt.Sprintf("{ \"message\": \"%s\" }", lowerCasedName))
		}

		we := waggy.WaggyError{
			Title:    "Internal server error",
			Detail:   "sorry, something went wrong!",
			Status:   0,
			Instance: "/",
		}

		handler := waggy.InitHandlerWithRoute("/lowercase", &flg).
			WithDefaultErrorResponse(we, http.StatusInternalServerError).
			MethodHandler(http.MethodPost, lowerCaserHandler)

		handler.ServeHTTP(w, r)
	})
}

func main() {}
