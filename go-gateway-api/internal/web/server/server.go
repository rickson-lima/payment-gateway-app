package server

import (
	"log"
	"net/http"

	"github.com/rickson-lima/payment-gateway-app/go-gateway-api/internal/service"
	"github.com/rickson-lima/payment-gateway-app/go-gateway-api/internal/web/handlers"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	port           string
}

func NewServer(accountService *service.AccountService, port string) *Server {
	srv := &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		port:           port,
	}

	srv.ConfigureRoutes()
	log.Println("Server up and running on port", port)

	return srv
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}
