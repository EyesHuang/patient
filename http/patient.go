package http

import (
	"encoding/json"
	"net/http"
)

// HandlerGetAll Return all patients
func (s *Server) HandlerGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode([]string{"Success"})
	}
}

// HandlerUpdateOrder Update a medical order
func (s *Server) HandlerUpdateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode([]string{"Success"})
	}
}
