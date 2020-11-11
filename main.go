package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Yo, HTTP!\n how you doin yeah?</h1>")
	})
	log.Fatal(http.ListenAndServe(":3001", nil))
}
