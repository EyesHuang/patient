package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Server is an HTTP server which embeds a chi router
type Server struct {
	router *chi.Mux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// NewServer creates and configures a new Server instance
func NewServer() *Server {
	s := Server{
		router: chi.NewRouter(),
	}
	s.routes()
	return &s
}
