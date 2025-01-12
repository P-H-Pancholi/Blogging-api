package models

type Post struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title" validate:"required,min=3"`
	Content string `json:"content" validate:"required"`
	Author  string `json:"author" validate:"required"`
}
