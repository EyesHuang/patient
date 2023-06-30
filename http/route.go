package http

func (s *Server) routes() {
	s.router.Get("/patients", s.HandlerGetAll())
	s.router.Put("/orders", s.HandlerUpdateOrder())
}
