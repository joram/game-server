package main

import (
	"fmt"
	"log"
	"net/http"
)



func handleRequests() {
	//http.HandleFunc("/", serveHomepage)
	http.HandleFunc("/pixels", servePixels)
	http.HandleFunc("/chunks", serveChunks)
	http.HandleFunc("/chat", chat)
	http.HandleFunc("/objects", objects)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/",  fs)
	log.Fatal(http.ListenAndServe(":2303", nil))
}

func main() {
	fmt.Println("starting server on port :2303")
	handleRequests()
}
