package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// this picks the index.html file to render in the current directory
	http.Handle("/", http.FileServer(http.Dir(".")))

	// default port is 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
