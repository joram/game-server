package main

import (
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/pixels", pixels)
	log.Fatal(http.ListenAndServe(":2303", nil))
}

func main() {
	handleRequests()
}
