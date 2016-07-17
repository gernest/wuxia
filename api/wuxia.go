package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Service struct {
	*mux.Router
}

func (s *Service) Todo(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("todo"))
}

func (s *Service) Init() {
	s.HandleFunc("/build", s.Todo)
	s.HandleFunc("/progress", s.Todo)
}
