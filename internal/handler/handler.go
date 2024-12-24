package handler

import (
    "fmt"
    "net/http"
)

type Handler struct{}

// コンストラクタ関数
// Go に new キーワードは存在しないので、慣習的にNewを定義するらしい
func New() *Handler {
    return &Handler{}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Go TODO Backend!")
}

func (h *Handler) TodoList(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "TODO List Endpoint")
}