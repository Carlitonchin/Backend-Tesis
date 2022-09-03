package handler

import (
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
	account := c.R.Group("api/account")

	account.POST("/signup", h.signUp)
	account.GET("/me", middleware.AuthUser(h.TokenService), h.me)
	account.POST("/signin", h.signin)
	account.POST("/signout", middleware.AuthUser(h.TokenService), h.signout)
	account.POST("/tokens", h.tokens)

	signedIn := c.R.Group("api")
	signedIn.Use(middleware.AuthUser(h.TokenService))

	signedIn.GET("/roles", middleware.OnlyAdmin(), h.getAllRoles)
	signedIn.GET("/users", middleware.OnlyAdmin(), h.getAllUsers)
	signedIn.PUT("/users/update-role", middleware.OnlyAdmin(), h.updateUserRole)
	signedIn.POST("/add-question", middleware.OnlyStudent(), h.addQuestion)
}
