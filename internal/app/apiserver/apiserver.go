package apiserver

import (
	"github.com/dmaok/web-1.1/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}
	s.logger.Info("starting api server...")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	logrus.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.HandleHello())
}

func (s *APIServer) HandleHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := io.WriteString(writer, "Hello Golang!")
		if err != nil {
			s.logger.Error(err)
		}
	}
}

func (s APIServer) configureStore() error {
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}
