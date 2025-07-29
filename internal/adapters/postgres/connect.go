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
	if cfg.Server.App.Mode == "production" {
		connStr = fmt.Sprintf(
			"user=%s password=%s dbname=%s port=%s sslmode=%s host=%s",
			cfg.DB.Prod.User,
			cfg.DB.Prod.Password,
			cfg.DB.Prod.Name,
			cfg.DB.Prod.Port,
			cfg.DB.Prod.SSLMode,
			cfg.DB.Prod.Host,
		)
	}
	if cfg.Server.App.Mode == "development" {
		connStr = fmt.Sprintf(
			"user=%s password=%s dbname=%s port=%s sslmode=%s host=%s",
			cfg.DB.Dev.User,
			cfg.DB.Dev.Password,
			cfg.DB.Dev.Name,
			cfg.DB.Dev.Port,
			cfg.DB.Dev.SSLMode,
			cfg.DB.Dev.Host,
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
