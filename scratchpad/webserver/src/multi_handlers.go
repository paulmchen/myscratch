package main

import (
	"fmt"
	"net/http"
	"strings"
)

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "This is my homepage!")
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "John Smith"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	fmt.Printf("parts=%s\n", parts)
	name := parts[2]
	if name == "" {
		name = "John Smith"
	}
	fmt.Fprint(res, "Goodbye ", name, "!")
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye/", goodbye)
	port := "8080"
	fmt.Printf("The server is listening on port %s.", port)
	http.ListenAndServe(":"+port, nil)
}
