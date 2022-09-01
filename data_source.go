package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/service"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dataSource struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func seed(db *gorm.DB) {
	var role *model.Role

	err := db.First(&role).Error

	if err != nil {
		erra := db.Create(&model.Role{Name: os.Getenv("ROLE_ADMIN")}).Error
		errc := db.Create(&model.Role{Name: os.Getenv("ROLE_DEFAULT_WORKER")}).Error
		erre := db.Create(&model.Role{Name: os.Getenv("ROLE_DEFAULT_STUDENT")}).Error

		errf := db.First(&role, "name = ?", "Administrador").Error

		hashed_pass, errp := service.HashPass("administrador")

		erru := db.Create(&model.User{
			Email:    "admin@admin.com",
			Name:     "admin",
			Password: hashed_pass,
			RoleID:   role.ID,
		}).Error

		if erra != nil || errc != nil || erru != nil || errf != nil || errp != nil || erre != nil {
			log.Fatal("Error seeding databbase")
		}
	}
}

func init_db() (*dataSource, error) {
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_user, db_pass, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err == nil {
		db.AutoMigrate(&model.Role{})
		db.AutoMigrate(&model.User{})
		seed(db)
	} else {
		log.Fatalf("Failed connection to postgres database, error: %v", err)
	}

	redis_host := os.Getenv("REDIS_HOST")
	redis_port := os.Getenv("REDIS_PORT")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redis_host, redis_port),
		Password: "",
		DB:       0,
	})

	_, err = rdb.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalf("Failed connection to redis database, error: %v", err)
	}

	return &dataSource{
		DB:          db,
		RedisClient: rdb,
	}, err
}
