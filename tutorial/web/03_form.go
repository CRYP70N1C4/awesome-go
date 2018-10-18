package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //important
	if r.Method == "GET" {

		t, _ := template.ParseFiles("D:\\code\\github\\awesome-go\\web\\main\\login.html")
		t.Execute(w, nil)

	} else {
		fmt.Println(r.Form)
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if username == "root" && password == "123456" {
			fmt.Fprint(w, "login success")
		} else {
			fmt.Fprint(w, "login failed")
		}
	}
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9090", nil)
}
