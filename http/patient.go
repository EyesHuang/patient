package http

import (
	"net/http"
)

// HandlerGetAll Return all patients
func (s *Server) HandlerGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(r, w, []string{"Success"}, http.StatusOK)
	}
}

// HandlerUpdateOrder Update a medical order
func (s *Server) HandlerUpdateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(r, w, []string{"Success"}, http.StatusOK)
	}
}
