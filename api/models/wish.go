package models

import (
	"gorm.io/gorm"
)

type Wish struct {
	gorm.Model
	Title  string `form:"title"`
	Detail string `form:"detail"`
	UserID string `form:"user_id"`
}
