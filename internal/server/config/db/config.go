package db

import (
	"log"
	"os"

	"github.com/OpsOMI/S.L.A.M/internal/server/config/db/models"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Prod models.DatabaseConfig
	Dev  models.DatabaseConfig
}

func LoadAll(files ...string) *DatabaseConfig {
	for _, file := range files {
		if err := godotenv.Overload(file); err != nil {
			log.Fatalf("error loading env file %s: %v", file, err)
		}
	}

	return &DatabaseConfig{
		Prod: models.DatabaseConfig{
			Driver:   os.Getenv("PROD_DB_DRIVER"),
			Host:     os.Getenv("PROD_DB_HOST"),
			SSLMode:  os.Getenv("PROD_DB_SSLMODE"),
			Timezone: os.Getenv("PROD_DB_TIMEZONE"),
			User:     os.Getenv("PROD_DB_USER"),
			Password: os.Getenv("PROD_DB_PASSWORD"),
			Port:     os.Getenv("PROD_DB_PORT"),
			Name:     os.Getenv("PROD_DB_NAME"),
		},
		Dev: models.DatabaseConfig{
			Driver:   os.Getenv("DEV_DB_DRIVER"),
			Host:     os.Getenv("DEV_DB_HOST"),
			SSLMode:  os.Getenv("DEV_DB_SSLMODE"),
			Timezone: os.Getenv("DEV_DB_TIMEZONE"),
			User:     os.Getenv("DEV_DB_USER"),
			Password: os.Getenv("DEV_DB_PASSWORD"),
			Port:     os.Getenv("DEV_DB_PORT"),
			Name:     os.Getenv("DEV_DB_NAME"),
		},
	}
}
