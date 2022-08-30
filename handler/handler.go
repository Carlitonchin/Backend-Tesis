package handler

import (
	"net/http"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService  model.UserService
	TokenService model.TokenService
}

type Config struct {
	R            *gin.Engine
	UserService  model.UserService
	TokenService model.TokenService
}

func NewHandler(c *Config) {
	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
	}

	g := c.R.Group("api/account")

	g.GET("/", h.Index)
	g.POST("/signup", h.SignUp)
	g.GET("/me", h.Me)
}

func (s *Handler) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
