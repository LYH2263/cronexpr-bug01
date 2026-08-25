package planner

import (
	"embed"
	"net/http"
)

//go:embed static/*
var staticFS embed.FS

// Server serves planner UI and API.
type Server struct {
	Handler http.Handler
}

func New(h http.Handler) *Server {
	if h == nil {
		h = http.NotFoundHandler()
	}
	mux := http.NewServeMux()
	mux.Handle("/api/", h)
	mux.Handle("/", http.FileServer(http.FS(staticFS)))
	return &Server{Handler: mux}
}
