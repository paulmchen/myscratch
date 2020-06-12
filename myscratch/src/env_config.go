package main

import (
	"fmt"
	"net/http"
	"os"
)

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}

func main() {
	http.HandleFunc("/", homePage)
	fmt.Printf("Port: %s\n", os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}