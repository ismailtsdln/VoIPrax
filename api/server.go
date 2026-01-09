package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ismailtsdln/VoIPrax/internal/logger"
)

// Server represents the REST API server
type Server struct {
	logger *logger.Logger
	router *gin.Engine
}

// NewServer creates a new REST API server instance
func NewServer(log *logger.Logger) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	s := &Server{
		logger: log,
		router: router,
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Add more API endpoints here (e.g., /fuzz, /exploit)
}

// Run starts the REST API server
func (s *Server) Run(addr string) error {
	s.logger.Info().Str("addr", addr).Msg("REST API server starting")
	return s.router.Run(addr)
}
