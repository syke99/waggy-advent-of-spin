package main

import (
	"encoding/json"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/syke99/waggy"
)

func init() {
	hello := map[string]string{
		"message": "Hello, world!",
	}

	helloBytes, _ := json.Marshal(hello)

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		waggy.WriteDefaultResponse(w, r)
	}

	handler := waggy.InitHandler(nil).
		WithDefaultResponse(helloBytes).
		MethodHandler(http.MethodGet, helloHandler)

	spinhttp.Handle(handler.ServeHTTP)
}

func main() {}
