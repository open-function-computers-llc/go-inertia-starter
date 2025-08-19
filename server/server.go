package server

import (
	"io/fs"
	"net/http"
	"os"
	"strconv"

	"github.com/petaki/inertia-go"
	"github.com/sirupsen/logrus"
)

type server struct {
	port           int
	logger         *logrus.Logger
	router         *http.ServeMux
	distFS         fs.FS
	inertiaManager *inertia.Inertia
}

func NewServer(fileSystem fs.FS, url string) server {
	portAsInt, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	s := server{
		port:           portAsInt,
		logger:         logrus.New(),
		router:         http.NewServeMux(),
		distFS:         fileSystem,
		inertiaManager: inertia.NewWithFS(url, "dist/index.html", "", fileSystem),
	}

	// share the app environment with the frontend
	s.inertiaManager.Share("appEnvironment", os.Getenv("APP_ENV"))

	s.inertiaManager.ShareFunc("assetPath", s.assetPath)

	s.logger.Info("about to bind routes!")
	s.bindRoutes()

	return s
}

func (s *server) Serve() error {
	s.logger.Info("now serving on port: " + strconv.Itoa(s.port))

	return http.ListenAndServe(":"+strconv.Itoa(s.port), s.router)
}
