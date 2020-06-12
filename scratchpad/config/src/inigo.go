package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, my name is John Smith.")
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Server is listening port 4000...")
	http.ListenAndServe("localhost:4000", nil)
}
