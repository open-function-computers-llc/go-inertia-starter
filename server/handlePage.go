package server

import "net/http"

func (s *server) handlePage(pageName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.inertiaManager.Render(w, r, pageName, nil)
	}
}
