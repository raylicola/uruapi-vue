package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Detail        string `form:"detail" binding:"required"`
	Rate          int    `form:"rate" binding:"required"`
	ReviewerID    string `form:"reviewer_id" binding:"required"`
	RevieweeID    string `form:"reviewee_id" binding:"required"`
}
