package apiserver

import (
	"encoding/json"
	"errors"
	"github.com/dmaok/web-1.1/internal/app/model"
	"github.com/dmaok/web-1.1/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	incorrectEmailOrPassword = errors.New("incorrect email or password")
	sessionName              = "web_1_session"
)

type server struct {
	store        store.Store
	router       *mux.Router
	logger       *logrus.Logger
	sessionStore sessions.Store
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		store:        store,
		logger:       logrus.New(),
		router:       mux.NewRouter(),
		sessionStore: sessionStore,
	}

	s.configureRouter()
	s.logger.Info("starting api server...")

	return s
}

func (s *server) ServeHTTP(writer http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(writer, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.HandleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.HandleSessionsCreate()).Methods("POST")
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) HandleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		user := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.store.User().Create(user); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		user.Sanitize()
		s.respond(w, r, http.StatusCreated, user)
	}
}

func (s *server) HandleSessionsCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, incorrectEmailOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}
