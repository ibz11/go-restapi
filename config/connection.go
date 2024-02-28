package config

import (
	"github.com/joho/godotenv"
	"os"
	"gorm.io/gorm"
	//"gorm.io/driver/postgres"
	"github.com/ibz11/go-restapi.git/storage"
	"github.com/ibz11/go-restapi.git/models"
	"log"
)

func ConnectDB() (*gorm.DB, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
		return nil, err
	}

	err = models.MigrateUsers(db)
	if err != nil {
		log.Fatal("could not migrate DB")
		return nil, err
	}

	return db, nil
}