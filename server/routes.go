package server

import "net/http"

func (s *server) bindRoutes() {
	routes := map[string]http.HandlerFunc{
		"GET /": s.handlePage("Index"),
	}

	for path, handler := range routes {
		s.router.Handle(path, s.inertiaManager.Middleware(s.LogRequest(handler)))
	}

	s.router.Handle("GET /dist/", http.FileServer(http.FS(s.distFS)))
}
