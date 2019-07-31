package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	public := http.StripPrefix("/public", http.FileServer(http.Dir("public")))

	http.Handle("/public/", public)
	http.HandleFunc("/", indexFunc)
	http.HandleFunc("/hello", helloHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
