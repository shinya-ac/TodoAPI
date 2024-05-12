package router

import (
	"net/http"

	"github.com/shinya-ac/TodoAPI/internal/adapter/handler"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", handler.HealthCheckHandler)
	return mux
}
