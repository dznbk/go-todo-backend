package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

	"github.com/dznbk/go-todo-backend/internal/handler"
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

    h := handler.New()

    http.HandleFunc("/", h.HealthCheck)
    http.HandleFunc("/todos", h.TodoList)

    fmt.Printf("Starting server on :%s\n", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}