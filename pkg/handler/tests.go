package handler

import (
	vpr "example"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTestHandler(c *gin.Context) {
	var body vpr.Test

	if err := c.BindJSON(&body); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.CreateTestService(body)

	if err != nil {
		newErrorResponse(c, err.Status, err.Message)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully created a test",
	})
}

func (h *Handler) GetTestsHandler(c *gin.Context) {
	tests, err := h.services.GetTestsService()

	if err != nil {
		newErrorResponse(c, err.Status, err.Message)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"tests": tests,
	})
}