package main

import (
	"log"
	"net/http"

	"github.com/gophergala/melted_brains/http_handler"
)

func main() {
	http.HandleFunc("/game/", http_handler.GameHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
