package handler

import (
	"time"

	"github.com/Carlitonchin/Backend-Tesis/handler/middleware"
	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService     model.UserService
	TokenService    model.TokenService
	RoleService     model.RoleService
	AreaService     model.AreaService
	QuestionService model.QuestionService
	ChatService     model.ChatService
}

type Config struct {
	R               *gin.Engine
	UserService     model.UserService
	TokenService    model.TokenService
	RoleService     model.RoleService
	AreaService     model.AreaService
	TimeOut         time.Duration
	QuestionService model.QuestionService
	ChatService     model.ChatService
}

func NewHandler(c *Config) {
	h := &Handler{
		UserService:     c.UserService,
		TokenService:    c.TokenService,
		RoleService:     c.RoleService,
		AreaService:     c.AreaService,
		QuestionService: c.QuestionService,
		ChatService:     c.ChatService,
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
	signedIn.GET("/questions", middleware.OnlyRoles([]string{"ROLE_ADMIN", "ROLE_DEFAULT_WORKER"}), h.getUnclasifiedQuestions)
	signedIn.GET("/questions/taken", middleware.NotTheseRoles([]string{"ROLE_DEFAULT_STUDENT"}), h.getTakenQuestions)
	signedIn.GET("/questions/taken/:id", middleware.NotTheseRoles([]string{"ROLE_DEFAULT_STUDENT"}), h.getTakenQuestionById)
	signedIn.GET("/questions/:status", h.getQuestionByStatus)
	signedIn.GET("/my-questions", middleware.OnlyRoles([]string{"ROLE_DEFAULT_STUDENT"}), h.getMyQuestions)
	signedIn.POST("/questions/add", middleware.OnlyRoles([]string{"ROLE_DEFAULT_STUDENT"}), h.addQuestion)
	signedIn.PUT("/questions/clasify", middleware.OnlyRoles([]string{"ROLE_ADMIN", "ROLE_DEFAULT_WORKER"}), h.clasifyQuestion)
	signedIn.PUT("/questions/take", middleware.OnlyRoles([]string{"ROLE_ADMIN", "ROLE_SPECIALIST_LEVEL1", "ROLE_SPECIALIST_LEVEL2"}), h.takeQuestion)
	signedIn.PUT("/questions/response", h.responseQuestion)
	signedIn.PUT("/questions/up-level", h.upLevel)
	signedIn.PUT("/questions/up-to-admin", h.upToAdmin)
	signedIn.GET("/areas", middleware.OnlyRoles([]string{"ROLE_ADMIN", "ROLE_DEFAULT_WORKER"}), h.get_areas)
	signedIn.POST("areas/add", middleware.OnlyRoles([]string{"ROLE_ADMIN"}), h.addArea)
	signedIn.GET("/chat/:question_id", h.getMessagesChat)
}
