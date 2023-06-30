package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	patient "manage-patinets"

	"github.com/go-chi/chi/v5"
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

		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)
		if err != nil {
			s.respond(r, w, err, http.StatusBadRequest)
			return
		}

		order, err := s.patientRepo.GetOrderByID(ctx, int32(id))
		if err != nil {
			s.respond(r, w, err, http.StatusInternalServerError)
			return
		}

		var req *patient.Order
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.respond(r, w, nil, http.StatusBadRequest)
			return
		}

		if err := s.patientRepo.UpdateOrder(ctx, req); err != nil {
			s.respond(r, w, err, http.StatusInternalServerError)
			return
		}

		s.respond(r, w, order, http.StatusOK)
	}
}
