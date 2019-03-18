package app

import "github.com/zeyneloz/sample-go-rest-api/internal/pkg/models"

// MigrateModels of all apps.
func MigrateModels() {
	models.Migrate()
}
