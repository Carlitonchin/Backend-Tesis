package main

import (
	"log"
	"os"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/service"
	"gorm.io/gorm"
)

func seed(db *gorm.DB) {
	var role *model.Role

	err := db.First(&role).Error

	if err != nil {
		seedRoles(db)
		seedUsers(db)
	}
}

func seedRoles(db *gorm.DB) {

	db.AutoMigrate(&model.Role{})

	erra := db.Create(&model.Role{Name: os.Getenv("ROLE_ADMIN")}).Error
	errc := db.Create(&model.Role{Name: os.Getenv("ROLE_DEFAULT_WORKER")}).Error
	erre := db.Create(&model.Role{Name: os.Getenv("ROLE_DEFAULT_STUDENT")}).Error

	if erra != nil || errc != nil || erre != nil {
		log.Fatal("Error at seed data for roles")
	}
}

func seedUsers(db *gorm.DB) {

	db.AutoMigrate(&model.User{})
	var role *model.Role

	errf := db.First(&role, "name = ?", os.Getenv("ROLE_ADMIN")).Error
	hashed_pass, errp := service.HashPass("administrador")

	erru := db.Create(&model.User{
		Email:    "admin@admin.com",
		Name:     "admin",
		Password: hashed_pass,
		RoleID:   role.ID,
	}).Error

	if errf != nil || errp != nil || erru != nil {
		log.Fatal("Error seeding users at db")
	}

}
