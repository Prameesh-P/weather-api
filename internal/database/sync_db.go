package database

import "github.com/Pramessh-P/weather-api/internal/model"

func SyncDatabase() {
	DB.AutoMigrate(
		&model.User{},
		&model.UserActivity{},
	)
}