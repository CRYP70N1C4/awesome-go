package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/test" {
		fmt.Fprint(w, "hello world")
		return
	}
	http.NotFound(w, r)
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":8888", mux)
}
