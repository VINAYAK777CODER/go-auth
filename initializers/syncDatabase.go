package initializers

import "github.com/VINAYAK777CODER/go-auth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}