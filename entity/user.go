package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Phone     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
