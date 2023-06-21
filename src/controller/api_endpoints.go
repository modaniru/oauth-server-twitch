package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) generalFollows(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"message": "ok",
	})
}

func (h *Handler) getUser(c *gin.Context) {
	variable, _ := c.Get("id")
	id := variable.(int)
	response, err := h.service.IUserService.GetById(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
