package fiber

import (
	"bank_service/app"
	"bank_service/config"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app    *fiber.App
	config config.Config
	router *Router
}

func NewServer(application app.App, cfg config.Config) *Server {
	server := &Server{
		app:    fiber.New(),
		config: cfg,
		router: NewRouter(application),
	}

	server.router.SetupRoutes(server.app)

	return server
}

func (s *Server) Start() error {
	httpPort := s.config.Server.HTTPPort
	if httpPort == 0 {
		httpPort = 51021
	}

	return s.app.Listen(fmt.Sprintf(":%d", httpPort))
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
