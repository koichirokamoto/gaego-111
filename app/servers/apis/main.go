package main

import (
	"fmt"
	"gaego-111/app/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/apis/hello", handlers.Hello)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defauting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
