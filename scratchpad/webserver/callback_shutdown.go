package main

import (
	"fmt"
	"net/http"
	"os"
)

func shutdown(res http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	fmt.Fprint(res, "This is my homepage!")
}

func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homePage)
	port := "8080"
	fmt.Printf("Server is listening on port %s.\n", port)
	http.ListenAndServe(":"+port, nil)
}
