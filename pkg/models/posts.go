package models

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" validate:"required,min=3"`
	Content   string    `json:"content" validate:"required"`
	Category  string    `json:"category" validate:"required"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
