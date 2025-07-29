package managment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ManagementConfig struct {
	Username string
	Password string
}

var defaultPath = "./env/real/.env.managment"

func LoadAll(
	envPath string,
) *ManagementConfig {
	if envPath == "" {
		envPath = defaultPath
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("error loading env file %s: %v", envPath, err)
	}

	return &ManagementConfig{
		Username: os.Getenv("MANAGEMENT_USERNAME"),
		Password: os.Getenv("MANAGEMENT_PASSWORD"),
	}
}
