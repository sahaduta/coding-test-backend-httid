package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
