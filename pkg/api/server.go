package api

import (
	"github.com/ajalck/spine/pkg/api/handler/interfaces"
	"github.com/ajalck/spine/pkg/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	port   string
}

func NewServerHTTP(
	cfg *config.Config,
	authHandler interfaces.AuthHandler,

) *Server {
	engine := gin.New()
	engine.Use(gin.Logger())

	api := engine.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/signup", authHandler.SignUp)
	}

	return &Server{
		engine: engine,
		port:   cfg.Port,
	}
}

func (s *Server) Start() error {
	return s.engine.Run(":" + s.port)
}
