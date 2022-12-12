package main

import (
	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/syke99/waggy"
	"net/http"
)

var flg waggy.FullServer

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		// "/" 404 handler is defined in 404_handler.go
		handler404 := Handler404(flg)

		// the handler is defined in hello_handler.go
		helloHandler := HelloHandler(flg)

		// initialize a router and route the handlers to the correct routes
		router := waggy.InitRouter(&flg).
			Handle("/", handler404).
			Handle("/hello/{name}", helloHandler)

		// pass the incoming request to router.ServeHTTP()
		router.ServeHTTP(w, r)
	})
}

func main() {}
