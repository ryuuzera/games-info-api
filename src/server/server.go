package server

import (
	"games-info-api/src/routes"

	"github.com/gin-gonic/gin"
	// "games-info-api/src/routes"
)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	
	router := gin.Default();

	return &Server{router: router}
}

func (s *Server) Run() {
	routes.SetRoutes(s.router)
	s.router.Run(":8080")
}

func (s *Server) Router() *gin.Engine {
	return s.router
}