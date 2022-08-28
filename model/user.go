package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Worker   bool   `json:"worker"`
}
