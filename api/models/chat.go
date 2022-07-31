package models

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Content string `form:"content" binding:"required"`
	UserID  string `form:"user_id" binding:"required"`
	ItemID  uint   `form:"item_id" binding:"required"`
}
