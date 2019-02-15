package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi there, I love %s!<h1>", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Now server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
