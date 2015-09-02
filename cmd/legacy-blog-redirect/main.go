package main

import (
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Redirecting request for", r.RequestURI)
	http.Redirect(w, r, os.Getenv("HOST")+r.RequestURI, http.StatusMovedPermanently)
}

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if host == "" {
		log.Fatal("$HOST must be set")
	}

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
