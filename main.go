package main

import (
	"log"

	"github.com/Carlitonchin/Backend-Tesis/handler"
	"github.com/Carlitonchin/Backend-Tesis/repository"
	"github.com/Carlitonchin/Backend-Tesis/service"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("starting server ...")
	router := gin.Default()
	db, err := repository.Connect()

	if err != nil {
		log.Fatalf("Error when connecting with db, error:%v", err)
	}

	user_repo := repository.NewUserRepository(db)
	us_config := &service.USConfig{UserRepository: user_repo}

	user_serv := service.NewUserService(us_config)

	c := handler.Config{
		R:           router,
		UserService: user_serv,
	}

	handler.NewHandler(&c)
	router.Run("localhost:8888")
}
