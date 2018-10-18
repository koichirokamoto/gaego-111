package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		log.Println(f.Name())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request is coming")
		http.ServeFile(w, r, "./index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defauting to port %s", port)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
