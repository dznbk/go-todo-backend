package main

import (
    "fmt"
	"log"
    "net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Go TODO Backend!")
    })

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "TODO List Endpoint")
	})

	fmt.Printf("Starting server on :%s\n", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
