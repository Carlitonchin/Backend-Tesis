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
	AreaService  model.AreaService
}

type Config struct {
	R            *gin.Engine
	UserService  model.UserService
	TokenService model.TokenService
	RoleService  model.RoleService
	AreaService  model.AreaService
	TimeOut      time.Duration
}

func NewHandler(c *Config) {
	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
		RoleService:  c.RoleService,
		AreaService:  c.AreaService,
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

	signedIn.GET("/roles", middleware.OnlyRoles([]string{"ROLE_ADMIN"}), h.getAllRoles)
	signedIn.GET("/users", middleware.OnlyRoles([]string{"ROLE_ADMIN"}), h.getAllUsers)
	signedIn.PUT("/users/update-role", middleware.OnlyRoles([]string{"ROLE_ADMIN"}), h.updateUserRole)
	signedIn.PUT("/users/update-area", middleware.OnlyRoles([]string{"ROLE_ADMIN"}), h.updateUserArea)
	signedIn.POST("/questions/add", middleware.OnlyRoles([]string{"ROLE_DEFAULT_STUDENT"}), h.addQuestion)
	signedIn.PUT("/questions/clasify", middleware.OnlyRoles([]string{"ROLE_ADMIN", "ROLE_DEFAULT_WORKER"}), h.clasifyQuestion)
	signedIn.POST("areas/add", middleware.OnlyRoles([]string{"ROLE_ADMIN"}), h.addArea)
}
