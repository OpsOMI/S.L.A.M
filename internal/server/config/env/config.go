package env

import (
	"log"
	"os"

	"github.com/OpsOMI/S.L.A.M/internal/server/config/env/models"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Managment models.Managment
	Jwt       models.Jwt
}

var defaultPath = "./env/real/.env"

func LoadAll(
	envPath string,
) *EnvConfig {
	if envPath == "" {
		envPath = defaultPath
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("error loading env file %s: %v", envPath, err)
	}

	return &EnvConfig{
		Managment: models.Managment{
			Username: os.Getenv("MANAGEMENT_USERNAME"),
			Password: os.Getenv("MANAGEMENT_PASSWORD"),
		},
		Jwt: models.Jwt{
			Issuer: os.Getenv("JWT_ISSUER"),
			Secret: os.Getenv("JWT_SECRET"),
		},
	}
}
