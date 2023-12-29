package server

import "net/http"

func (s *server) LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info(r.Method+": ", r.URL.Path)

		next(w, r)
	}
}
