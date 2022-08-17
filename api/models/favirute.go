package models

import (
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserID    string `form:"user_id" binding:"required"`
	ItemID    uint   `form:"item_id" binding:"required"`
	Title     string
	Price     int
	Img       string
}
