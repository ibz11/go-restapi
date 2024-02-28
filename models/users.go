package models

import (
	"gorm.io/gorm"
)
type Users struct {
	gorm.Model
	FirstName    string   `json:"firstname"`
	Lastname     string   `json:"lastname"`
	Email        string   `json:"email"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	return err
}
