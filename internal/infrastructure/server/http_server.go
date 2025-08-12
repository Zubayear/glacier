package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"glacier/internal/presentation/http"
)

type Server struct {
	router      *chi.Mux
	userHandler *http.UserHandler
}

func NewServer(userHandler *http.UserHandler) *Server {
	return &Server{
		router:      chi.NewRouter(),
		userHandler: userHandler,
	}
}

func (s *Server) SetupRoutes() {
	s.router.Use(middleware.Recoverer)
	s.router.Post("/users", s.userHandler.CreateUser)
}

func (s *Server) GetRouter() *chi.Mux {
	return s.router
}
