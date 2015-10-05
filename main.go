package main

import (
	"log"
	"net/http"
	"os"

	// "chat"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// websocket server
	server := NewServer("/ws")
	go server.Listen()

	// static files
	http.Handle("/", http.FileServer(http.Dir("public")))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
