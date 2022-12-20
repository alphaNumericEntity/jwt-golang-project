package initializers

import "github.com/alphanumericentity/jwt-example/models"

func SyncDatabase() {
	// Auto Migrate
	DB.AutoMigrate(&models.User{})
}
