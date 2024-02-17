package entity

import (
	"time"

	"gorm.io/gorm"
)

type CustomPage struct {
	Id        uint `gorm:"primaryKey"`
	CustomUrl string
	Content   string
	UserId    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
