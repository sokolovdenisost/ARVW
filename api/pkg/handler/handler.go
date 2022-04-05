package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sokolovdenisost/VPR/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.GET("/", h.GetAuthHandler)
		auth.POST("sign-up", h.SignUpHandler)
		auth.POST("sign-in", h.SignInHandler)
	}

	return router
}
