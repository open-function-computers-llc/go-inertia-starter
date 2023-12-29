package server

import (
	"io/fs"
	"net/http"
	"strconv"

	"github.com/petaki/inertia-go"
	"github.com/sirupsen/logrus"
)

type server struct {
	port           int
	logger         *logrus.Logger
	router         *http.ServeMux
	inertiaManager *inertia.Inertia
	distFS         fs.FS
}

func New(port int, url string, fs fs.FS) (server, error) {
	s := server{
		logger:         logrus.New(),
		port:           port,
		router:         http.NewServeMux(),
		inertiaManager: inertia.New(url, "./server/index.html", ""),
		distFS:         fs,
	}

	s.bindRoutes()

	return s, nil
}

func (s *server) Serve() error {
	return http.ListenAndServe(":"+strconv.Itoa(s.port), s.router)
}
