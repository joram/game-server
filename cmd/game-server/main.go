package main

import (
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/", serveHomepage)
	http.HandleFunc("/pixels", servePixels)
	http.HandleFunc("/chunks", serveChunks)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Fatal(http.ListenAndServe(":2303", nil))
}

func main() {
	handleRequests()
}
