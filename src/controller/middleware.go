package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
)

func (h *Handler) authMiddleware(c *gin.Context) {
	token := c.GetHeader(authHeader)
	id, err := h.service.IUserService.ValidateJwtToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": err.Error(),
		})
	}
	c.Set("id", id)
}
