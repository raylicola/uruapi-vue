package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	PurchaserID  string  `form:"purchaser_id" binding:"required"`
	SellerID     string  `form:"seller_id" binding:"required"`
	ItemID       uint    `form:"item_id" binding:"required"`
}