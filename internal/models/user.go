package models

import "gorm.io/gorm"


type User struct {
	gorm.Model
	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type RegisterRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
