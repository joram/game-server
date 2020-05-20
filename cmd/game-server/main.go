package main

import (
	"fmt"
	"github.com/joram/game-server/game"
	"log"
	"net/http"
)



func handleRequests() {
	//http.HandleFunc("/", serveHomepage)
	http.HandleFunc("/pixels", game.ServePixels)
	http.HandleFunc("/chunks", game.ServeChunks)
	http.HandleFunc("/chat", game.ServeChat)
	http.HandleFunc("/objects", game.ServeObjects)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static", http.StripPrefix("/static/", fs))
	http.Handle("/",  fs)

	log.Fatal(http.ListenAndServe(":2303", nil))
}

func main() {
	fmt.Println("starting server on port :2303")
	handleRequests()
}
