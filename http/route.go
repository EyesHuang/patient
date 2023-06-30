package http

func (s *Server) routes(p ...Pinger) {
	s.router.Get("/healthcheck", s.handleHealthCheck(p...))
	s.router.Get("/patients", s.HandlerGetAll())
	s.router.Put("/orders/{id}", s.HandlerUpdateOrder())
}
