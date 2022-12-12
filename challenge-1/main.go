package main

import (
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/syke99/waggy"
)

var flg waggy.FullServer

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		handler := HelloHandler(flg)

		// pass the incoming request into the handler initialized above
		handler.ServeHTTP(w, r)
	})
}

func main() {}
