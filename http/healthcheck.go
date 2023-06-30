package http

import (
	"context"
	"net/http"
	"sync"
)

// Pinger is an interface that represents any type that can respond to a PingContext.
type Pinger interface {
	PingContext(ctx context.Context) error
}

func (s *Server) handleHealthCheck(ps ...Pinger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errCh := make(chan error, len(ps))
		var wg sync.WaitGroup

		for _, p := range ps {
			wg.Add(1)
			go func(p Pinger) {
				defer wg.Done()
				if err := p.PingContext(r.Context()); err != nil {
					errCh <- err
				}
			}(p)
		}

		wg.Wait()
		close(errCh)

		if len(errCh) > 0 {
			s.respond(r, w, nil, http.StatusInternalServerError)
			return
		}

		s.respond(r, w, nil, http.StatusOK)
	}
}
