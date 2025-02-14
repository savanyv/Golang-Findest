package app

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/savanyv/Golang-Findest/internal/config"
	"github.com/savanyv/Golang-Findest/internal/config/database"
	"github.com/savanyv/Golang-Findest/internal/delivery/routes"
	"github.com/savanyv/Golang-Findest/internal/middlewares"
)

type Server struct {
	e *echo.Echo
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		e: echo.New(),
		config: config,
	}
}

func (s *Server) Run() error {
	// Setup Database
	_, err := database.ConnectDB(*s.config)
	if err != nil {
		log.Println("failed to connect database")
	}

	// Setup Middlewares
	s.e.Use(middlewares.CORSMiddleware)
	s.e.Use(middlewares.MethodValidationMiddleware)

	// Setup Routes
	routes.RegisterRoutes(s.e)

	// Start Server
	if err := s.e.Start(":7000"); err != nil {
		log.Println("failed to start server")
		return err
	}

	log.Println("Server started on port 7000")
	return nil
}