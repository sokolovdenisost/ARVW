package handler

import (
	vpr "example"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetResultsHandler(c *gin.Context) {
	id := c.Param("id")

	results, err := h.services.Results.GetResultsService(id)

	if err != nil {
		newErrorResponse(c, err.Status, err.Message)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"results": results,
	})
}

func (h *Handler) CreateResultHandler(c *gin.Context) {
	var body vpr.Result

	if errJSON := c.BindJSON(&body); errJSON != nil {
		newErrorResponse(c, http.StatusInternalServerError, errJSON.Error())
		return
	}

	err := h.services.Results.CreateResultService(body)

	if err != nil {
		newErrorResponse(c, err.Status, err.Message)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully completed the test",
	})
}
