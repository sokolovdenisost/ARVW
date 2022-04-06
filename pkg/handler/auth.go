package handler

import (
	"fmt"
	"net/http"

	vpr "example"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUpHandler(c *gin.Context) {
	var body vpr.User

	if err := c.BindJSON(&body); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.CreateUserService(body)

	if err != nil {
		newErrorResponse(c, err.Status, err.Message)
		return
	}

	if user != nil {
		newErrorResponse(c, http.StatusBadRequest, "This email address exists")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success registration on this service",
	})
}

func (h *Handler) SignInHandler(c *gin.Context) {
	var body vpr.SignInBody

	if err := c.BindJSON(&body); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.GenerateTokenService(body)

	if err != nil {
		newErrorResponse(c, err.Status, err.Message)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) GetAuthHandler(c *gin.Context) {
	id, err := h.GetAuth(c)

	if err != nil {
		newErrorResponse(c, err.Status, err.Message)
		return
	}

	user, err := h.services.Authorization.GetUserByIdService(fmt.Sprint(id))

	if user == nil {
		newErrorResponse(c, http.StatusNotFound, "Not found")
		return
	}

	if err != nil {
		newErrorResponse(c, err.Status, err.Message)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
