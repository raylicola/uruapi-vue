package models

import (
	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	ItemID  uint   `form:"item_id" binding:"required"`
	WishID  uint   `form:"wish_id" binding:"required"`
}
