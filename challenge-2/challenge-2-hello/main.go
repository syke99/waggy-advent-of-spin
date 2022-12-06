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

var flg waggy.FullCGI

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		helloHandler := func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Content-Type", "text/plain")
				fmt.Fprintln(w, "resource not found")
				return
			}

			name := waggy.Vars(r)["name"]

			helloElfJSON := "{ \"message\":\"Hello, world\" }"

			if name != "" {
				lowerCaseName := ""

				resp, err := spinhttp.Post("127.0.0.1:3001/lowercase", "application/json", bytes.NewBufferString(fmt.Sprintf("{ \"message\": %s", name)))
				if err != nil {
					waggy.WriteDefaultErrorResponse(w, r)
				}
				defer resp.Body.Close()

				body := make([]byte, 0)

				if resp != nil {
					body, err = ioutil.ReadAll(resp.Body)
				}

				if err != nil {
					waggy.WriteDefaultErrorResponse(w, r)
				}

				splitBody := strings.Split(string(body), " ")

				if len(splitBody) > 3 {
					lowerCaseName = splitBody[2:3][0]
				}

				if len(lowerCaseName) != 0 &&
					lowerCaseName[:1] == "\"" &&
					lowerCaseName[len(lowerCaseName)-1:] == "\"" {
					lowerCaseName = lowerCaseName[1 : len(lowerCaseName)-1]
				}

				helloElfJSON = fmt.Sprintf("{ \"message\":\"Hello, %s\" }", lowerCaseName)
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, helloElfJSON)
		}

		we := waggy.WaggyError{
			Title:    "Internal server error",
			Detail:   "sorry, something went wrong!",
			Status:   0,
			Instance: "/",
		}

		handler := waggy.InitHandlerWithRoute("/hello/{name}", &flg).
			WithDefaultErrorResponse(we, http.StatusInternalServerError).
			MethodHandler(http.MethodGet, helloHandler)

		handler.ServeHTTP(w, r)
	})
}

func main() {}
