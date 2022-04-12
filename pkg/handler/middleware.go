package handler

import (
	"net/http"
	"strings"

	vpr "example"

	"github.com/gin-gonic/gin"
)

const userCTX = "userID"

func SetError(status int, message string) *vpr.Error {
	return &vpr.Error{Status: status, Message: message}
}

func (h *Handler) userIdentity(c *gin.Context) {
	userID, err := h.GetAuth(c)

	if err != nil {
		newErrorResponse(c, err.Status, err.Message)
		return
	}

	c.Set(userCTX, userID)
}

func (h *Handler) GetAuth(c *gin.Context) (string, *vpr.Error) {
	header := c.GetHeader("Authorization")

	if header == "" {
		return "", SetError(http.StatusUnauthorized, "empty auth header")
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		return "", SetError(http.StatusUnauthorized, "invalid auth header")
	}

	userID, err := h.services.Authorization.ParseTokenService(headerParts[1])

	if err != nil {
		return "", SetError(err.Status, err.Message)
	}

	return userID, nil
}
