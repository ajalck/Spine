package api

import (
	"github.com/ajalck/spine/pkg/api/handler/interfaces"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServerHTTP(
	authHandler interfaces.AuthHandler,

) *Server {
	engine := gin.New()
	engine.Use(gin.Logger())

	//api:=engine.Group("/api/v1")

	return &Server{
		engine: engine,
	}
}

func (s *Server) Start() error {
	return s.engine.Run(":8080")
}
