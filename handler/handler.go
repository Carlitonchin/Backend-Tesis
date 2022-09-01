package handler

import (
	"net/http"
	"time"

	"github.com/Carlitonchin/Backend-Tesis/handler/middleware"
	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService  model.UserService
	TokenService model.TokenService
	RoleService  model.RoleService
}

type Config struct {
	R            *gin.Engine
	UserService  model.UserService
	TokenService model.TokenService
	RoleService  model.RoleService
	TimeOut      time.Duration
}

func NewHandler(c *Config) {
	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
		RoleService:  c.RoleService,
	}
	timeouterror := apperrors.NewError(apperrors.TimeOut, "El request demoró mucho en procesarse")
	c.R.Use(middleware.Timeout(c.TimeOut, timeouterror))
	g := c.R.Group("api/account")

	g.GET("/", h.index)
	g.POST("/signup", h.signUp)
	g.GET("/me", middleware.AuthUser(h.TokenService), h.me)
	g.POST("/signin", h.signin)
	g.POST("/signout", middleware.AuthUser(h.TokenService), h.signout)
	g.POST("/tokens", h.tokens)
	g.GET("/roles", h.getAllRoles)
}

func (s *Handler) index(ctx *gin.Context) {
	time.Sleep(10 * time.Second)
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
