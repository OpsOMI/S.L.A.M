package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/configs/server"
	_ "github.com/lib/pq"
)

func Connect(
	cfg server.ServerConfigs,
) (*sql.DB, error) {
	var connStr string
	if cfg.AppEnv == "production" {
		connStr = fmt.Sprintf(
			"user=%s password=%s dbname=%s port=%s sslmode=%s host=%s",
			cfg.Env.Prod.User,
			cfg.Env.Prod.Password,
			cfg.Env.Prod.Name,
			cfg.Env.Prod.Port,
			cfg.Env.Prod.SSLMode,
			cfg.Env.Prod.Host,
		)
	}
	if cfg.AppEnv == "development" {
		connStr = fmt.Sprintf(
			"user=%s password=%s dbname=%s port=%s sslmode=%s host=%s",
			cfg.Env.Dev.User,
			cfg.Env.Dev.Password,
			cfg.Env.Dev.Name,
			cfg.Env.Dev.Port,
			cfg.Env.Dev.SSLMode,
			cfg.Env.Dev.Host,
		)
	}

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	for range 5 {
		err = conn.Ping()
		if err == nil {
			return conn, nil
		}
		if strings.Contains(err.Error(), "database system is starting up") {
			log.Println("Database is starting up, retrying...")
			time.Sleep(2 * time.Second)
			continue
		}
		return nil, err
	}

	return nil, err
}
