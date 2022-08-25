package handler

import (
	"net/http"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService model.UserService
}

type Config struct {
	R           *gin.Engine
	UserService model.UserService
}

func NewHandler(c *Config) {
	h := &Handler{
		UserService: c.UserService,
	}

	g := c.R.Group("api/account")

	g.GET("/", h.Index)
}

func (s *Handler) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
