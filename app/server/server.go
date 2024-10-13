package server

import (
	"github.com/Fagan04/http-rest-api/app/config"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *config.Config
	logger *log.Logger
	router *mux.Router
}

func NewAPIServer(config *config.Config) *APIServer {
	return &APIServer{
		config: config,
		logger: log.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.ConfigureLogger(); err != nil {
		return err
	}
	s.logger.Info("Starting API server")
	s.ConfigureRouter()

	return http.ListenAndServe(s.config.BindAddress, s.router)

}

func (s *APIServer) ConfigureLogger() error {
	level, err := log.ParseLevel(s.config.LogLevel)
	if err != nil {
		s.logger.Fatal("Could not parse log level")
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) ConfigureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World!")
	}
}
