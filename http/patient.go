package http

import (
	"net/http"
)

// HandlerGetAll Return all patients
func (s *Server) HandlerGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		patients, err := s.patientRepo.GetAllPatients(ctx)
		if err != nil {
			s.respond(r, w, err, http.StatusInternalServerError)
		}

		s.respond(r, w, patients, http.StatusOK)
	}
}

// HandlerUpdateOrder Update a medical order
func (s *Server) HandlerUpdateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(r, w, []string{"Success"}, http.StatusOK)
	}
}
