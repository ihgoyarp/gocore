package router

import (
	"net/http"

	"goc/internal/handler"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.HealthHandler)
	mux.HandleFunc("/echo", handler.EchoHandler)

	return mux
}
