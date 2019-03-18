package models

// Author struct describes the author model for GORM.
type Author struct {
	ID       int    `gorm:"primary_key"`
	Name     string `gorm:"size:70"`
	LastName string `gorm:"size:70"`
}

// Book struct describes the book model for GORM.
type Book struct {
	ID          int    `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(150);unique_index"`
	AuthorID    int
	Author      Author
	CreatedByID int
	CreatedBy   User
}
