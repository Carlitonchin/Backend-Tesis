package main

import (
	"log"

	"github.com/Carlitonchin/Backend-Tesis/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("starting server ...")
	router := gin.Default()

	c := handler.Config{
		R: router,
	}

	handler.NewHandler(&c)
	router.Run("localhost:8888")
}
