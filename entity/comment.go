package entity

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	Id            uint `gorm:"primaryKey"`
	NewsArticleId uint
	Name          string
	Content       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
