package main

import (
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	targetPath := r.RequestURI

	if targetPath == "/" {
		targetPath = "/about/"
	}

	url := os.Getenv("HOST") + targetPath

	log.Println("Redirecting request:", "Path:", r.RequestURI, "Referrer:", r.Header["Referer"], "To:", url)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
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
