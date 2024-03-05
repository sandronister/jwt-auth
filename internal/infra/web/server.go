package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sandronister/jwt-auth/internal/handlers"
)

type WebServer struct {
	router  *gin.Engine
	webPort string
}

func NewWebServer(webPort string) *WebServer {
	return &WebServer{
		router:  gin.Default(),
		webPort: webPort,
	}
}

func (s *WebServer) AddRegisterHandler(handler *handlers.UserHandler) {
	public := s.router.Group("/api")

	public.POST("/register", handler.Register)
}

func (s *WebServer) Start() error {
	return s.router.Run(s.webPort)
}
