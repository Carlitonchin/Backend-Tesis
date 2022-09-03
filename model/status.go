package model

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Description string `json:"description" gorm:"unique;not null"`
}
