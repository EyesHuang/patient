package http

import (
	"encoding/json"
	"net/http"

	patient "manage-patinets"
)

// HandlerGetAll Return all patients
func (s *Server) HandlerGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		patients, err := s.patientRepo.GetAllPatients(ctx)
		if err != nil {
			s.respond(r, w, err, http.StatusInternalServerError)
			return
		}

		s.respond(r, w, patients, http.StatusOK)
	}
}

// HandlerUpdateOrder Update a medical order
func (s *Server) HandlerUpdateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req *patient.Order
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.respond(r, w, nil, http.StatusBadRequest)
			return
		}

		order, err := s.patientRepo.Update(ctx, req)
		if err != nil {
			s.respond(r, w, err, http.StatusInternalServerError)
			return
		}

		s.respond(r, w, order, http.StatusOK)
	}
}
