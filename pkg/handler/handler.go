package handler

import (
	"example/pkg/service"

	"github.com/gin-contrib/cors"
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
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))

	auth := router.Group("/auth")
	{
		auth.GET("/", h.GetAuthHandler)
		auth.POST("/sign-up", h.SignUpHandler)
		auth.POST("/sign-in", h.SignInHandler)
	}

	api := router.Group("/api", h.userIdentity)
	{
		tests := api.Group("/tests")
		{
			tests.GET("/", h.GetTestsHandler)
			tests.GET("/:id", h.GetTestByIdHandler)
			tests.GET("/:id/answers", h.GetTestByIdWithAnswersHandler)
			tests.POST("/create", h.CreateTestHandler)
		}

		results := api.Group("/results")
		{
			results.GET("/:id", h.GetResultsHandler)
			results.POST("/create", h.CreateResultHandler)
		}
	}

	return router
}
