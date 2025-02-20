package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title         string `json:"title" gorm:"not null"`
	Author        string `json:"author" gorm:"not null"`
	PublishedDate string `json:"publishedDate"`
	Genre         string `json:"genre"`
	Description   string `json:"description"`
	CoverImageUrl string `json:"coverImageUrl,omitempty"`
	ISBN          string `json:"isbn,omitempty"`
}

type BookResponse struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"publishedDate"`
	Genre         string `json:"genre"`
	Description   string `json:"description"`
	CoverImageUrl string `json:"coverImageUrl,omitempty"`
	ISBN          string `json:"isbn,omitempty"`
}
