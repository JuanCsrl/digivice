package main

import (
	"log"
	"net/http"

	"example.com/m/handlers"
)

func main() {
	http.HandleFunc("/", handlers.DigiviceHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
