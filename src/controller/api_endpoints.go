package controller

import "github.com/gin-gonic/gin"

func (h *Handler) generalFollows(c *gin.Context){
	c.JSON(200, map[string]interface{}{
		"message": "ok",
	})
}