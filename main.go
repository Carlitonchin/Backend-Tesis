package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("starting server ...")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	db, err := init_db()

	if err != nil {
		log.Fatalf("Error when connecting with db, error:%v", err)
	}

	router, err := inject(db)

	if err != nil {
		log.Fatal(err)
	}

	router.Run("localhost:8888")
}
