package models

import "gorm.io/gorm"


type Portfolio struct {
	gorm.Model
	Title string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	ImageURL string `json:"image_url"`
	ProjectURL string `json:"project_url"`
	UserID uint `json:"user_id"`
}
