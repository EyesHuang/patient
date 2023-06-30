package http

import (
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

		reqOrder, err := preProcessData(r)
		if err != nil {
			s.respond(r, w, err, http.StatusBadRequest)
			return
		}

		if _, err := s.patientRepo.GetOrderByID(ctx, int32(reqOrder.ID)); err != nil {
			s.respond(r, w, err, http.StatusInternalServerError)
			return
		}

		if err := s.patientRepo.UpdateOrder(ctx, reqOrder); err != nil {
			s.respond(r, w, err, http.StatusInternalServerError)
			return
		}

		s.respond(r, w, reqOrder, http.StatusOK)
	}
}

func preProcessData(r *http.Request) (*patient.Order, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		return nil, err
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return nil, err
	}

	message := r.Form.Get("message")

	order := &patient.Order{
		ID:      int(id),
		Message: message,
	}

	return order, err
}
