package server

import (
	"net/http"
)

func (s *server) bindRoutes() {
	openRoutes := map[string]http.HandlerFunc{
		"/about": s.handleAbout(),
		"/login": s.handleLogin(),
		"/":      s.handleIndex(),
	}

	for path, handler := range openRoutes {
		s.router.Handle(path, s.inertiaManager.Middleware(s.LogRequest(handler)))
	}

	// The following code could be uncommented, and wrapped in an additional middleware (Protect) to protect certain routes
	//
	// protectedRoutes := map[string]http.HandlerFunc{
	// 	"/protected": s.handleProtected(),
	// }

	// for path, handler := range protectedRoutes {
	// 	s.router.Handle(path, s.Protect(s.inertiaManager.Middleware(s.LogRequest(handler))))
	// }

	s.router.Handle("/dist/", http.FileServer(http.FS(s.distFS)))
}
