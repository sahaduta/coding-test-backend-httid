package entity

import (
	"time"

	"gorm.io/gorm"
)

type NewsArticle struct {
	Id         uint `gorm:"primaryKey"`
	CategoryId uint
	Content    string
	UserId     uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
