package main

import (
	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/syke99/waggy"
	"net/http"
)

var flg waggy.FullCGI

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		// the handler is defined in hello_handler.go
		handler := HelloHandler(flg)

		// pass the incoming request into the handler initialized above
		handler.ServeHTTP(w, r)
	})
}

func main() {}
