package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/modaniru/api-for-users/src/model"
)

func (h *Handler) signIn(c *gin.Context){
	var access model.AccessToken
	err := c.BindJSON(&access)
	if err != nil{
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}
	token, err := h.service.IUserService.Login(access.AccessToken)
	if err != nil{
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"jwt": token,
	})
}