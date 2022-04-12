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
		"status":  200,
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
		"status": 200,
		"tests":  tests,
	})
}

func (h *Handler) GetTestByIdHandler(c *gin.Context) {
	id := c.Param("id")

	test, err := h.services.GetTestByIdService(id, false)

	if err != nil {
		newErrorResponse(c, err.Status, err.Message)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"test":   test,
	})
}

func (h *Handler) GetTestByIdWithAnswersHandler(c *gin.Context) {
	id := c.Param("id")

	test, err := h.services.GetTestByIdService(id, true)

	if err != nil {
		newErrorResponse(c, err.Status, err.Message)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"test":   test,
	})
}
