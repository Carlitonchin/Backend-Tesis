package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Text      string `json:"text" gorm:"not null"`
	UserRefer uint   `json:"-"`
}
