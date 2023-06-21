package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/modaniru/api-for-users/src/service"
)

type Handler struct{
	service *service.Service
}

func NewHandler(service *service.Service) *Handler{
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRouters() http.Handler{
	engine := gin.New()
	auth := engine.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
	}
	api := engine.Group("/api")
	{
		api.GET("/general-follows", h.generalFollows)
	}
	return engine
}