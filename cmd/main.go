package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Go TODO Backend!")
    })

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "TODO List Endpoint")
	})

    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
