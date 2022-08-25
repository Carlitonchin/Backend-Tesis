package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("starting server ...")
	router := gin.Default()

	router.GET("api/account/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Hello": "World",
		})
	})

	router.Run("localhost:8888")
}
