package main

import (
	"log"
	"net/http"

	"./../gopher"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	http.HandleFunc("/", gopher.FirestoreLookup)

	log.Printf("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
