package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{}

	g := c.R.Group("api/account")

	g.GET("/", h.Index)
}

func (s *Handler) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
