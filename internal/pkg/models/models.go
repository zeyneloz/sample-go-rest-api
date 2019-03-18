package models

import "github.com/zeyneloz/sample-go-rest-api/internal/svc/db"

// Migrate User models.
func Migrate() {
	// Migrate the user model.
	db.Instance.AutoMigrate(&User{})
	db.Instance.AutoMigrate(&Author{})
	db.Instance.AutoMigrate(&Book{})
}
