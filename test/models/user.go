package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm: "unqueIdex;not null"`
	Password string  `gorm:"not null"`
}