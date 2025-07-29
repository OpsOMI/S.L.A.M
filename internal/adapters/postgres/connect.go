package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/server/config"
	_ "github.com/lib/pq"
)

func Connect(
	serverConfig config.Configs,
) (*sql.DB, error) {
	var connStr string
	if serverConfig.Server.App.Mode == "production" {
		connStr = fmt.Sprintf(
			"user=%s password=%s dbname=%s port=%s sslmode=%s host=%s",
			serverConfig.DB.Prod.User,
			serverConfig.DB.Prod.Password,
			serverConfig.DB.Prod.Name,
			serverConfig.DB.Prod.Port,
			serverConfig.DB.Prod.SSLMode,
			serverConfig.DB.Prod.Host,
		)
	}
	if serverConfig.Server.App.Mode == "development" {
		connStr = fmt.Sprintf(
			"user=%s password=%s dbname=%s port=%s sslmode=%s host=%s",
			serverConfig.DB.Dev.User,
			serverConfig.DB.Dev.Password,
			serverConfig.DB.Dev.Name,
			serverConfig.DB.Dev.Port,
			serverConfig.DB.Dev.SSLMode,
			serverConfig.DB.Dev.Host,
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
