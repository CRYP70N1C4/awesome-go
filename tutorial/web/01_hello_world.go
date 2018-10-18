package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("form", r.Form)
	fmt.Println("Path", r.URL.Path)
	fmt.Println("Scheme", r.URL.Scheme)
	fmt.Println("User", r.URL.User)
	for k, v := range r.Form {
		fmt.Println(k, " --> ", strings.Join(v,""))
	}
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/test", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
