package handler

import (
	"example/pkg/service"

	"github.com/gin-gonic/gin"
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

	api := router.Group("/api")
	{
		tests := api.Group("/tests")
		{
			tests.GET("/", h.GetTestsHandler)
			tests.GET("/:id", h.GetTestByIdHandler)
			tests.GET("/:id/answers", h.GetTestByIdWithAnswersHandler)
			tests.POST("/:id/answers", h.SendAnswersHandler)
			tests.POST("/create", h.CreateTestHandler)
		}
	}

	return router
}
