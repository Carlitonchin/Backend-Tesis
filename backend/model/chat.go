package model

import "gorm.io/gorm"

type MessageChat struct {
	gorm.Model
	Question uint   `json:"question_id"`
	User     uint   `json:"-"`
	Text     string `json:"text"`
	Readed   bool   `json:"-"`
	UserName string `gorm:"-" json:"user_name"`
}
