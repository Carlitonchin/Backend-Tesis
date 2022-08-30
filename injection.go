package main

import (
	"github.com/Carlitonchin/Backend-Tesis/handler"
	"github.com/Carlitonchin/Backend-Tesis/repository"
	"github.com/Carlitonchin/Backend-Tesis/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func inject(db *gorm.DB) (*gin.Engine, error) {

	router := gin.Default()
	user_repo := repository.NewUserRepository(db)
	us_config := &service.USConfig{UserRepository: user_repo}

	user_serv := service.NewUserService(us_config)

	c := handler.Config{
		R:           router,
		UserService: user_serv,
	}

	handler.NewHandler(&c)

	return router, nil
}
