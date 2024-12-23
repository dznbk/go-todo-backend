package handler

import (
    "fmt"
    "net/http"
)

type Handler struct{}

func New() *Handler {
    return &Handler{}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Go TODO Backend!")
}

func (h *Handler) TodoList(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "TODO List Endpoint")
}