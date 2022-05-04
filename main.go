package main

import (
	"log"
	"net/http"
)

// main function
func main() {
	http.HandleFunc("/", handle)
	log.Println("Listening on addr", 8080)
	http.ListenAndServe(":8080", nil)

}
