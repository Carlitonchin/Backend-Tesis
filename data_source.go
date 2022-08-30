package main

import (
	"fmt"
	"os"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init_db() (*gorm.DB, error) {
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_user, db_pass, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err == nil {
		db.AutoMigrate(&model.User{})
	}

	return db, err
}
