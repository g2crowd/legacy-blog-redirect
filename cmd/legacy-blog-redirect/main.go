package main

import (
	"log"
	"os"
	"net/http"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if host == "" {
		log.Fatal("$HOST must be set")
	}

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.Handle("/", http.RedirectHandler(host, http.StatusMovedPermanently))
	http.ListenAndServe(":"+port, nil)
}
