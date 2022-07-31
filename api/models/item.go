package models

import (
  "gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Title      string  `form:"title" binding:"required"`
	Detail     string  `form:"detail" binding:"required"`
	Price      int     `form:"price" binding:"required"`
	Img        string  `form:"img" binding:"required"`
	UserID     string  `form:"user_id" binding:"required"`
}