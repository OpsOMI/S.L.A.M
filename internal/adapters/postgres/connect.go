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
			serverConfig.Db.Prod.User,
			serverConfig.Db.Prod.Password,
			serverConfig.Db.Prod.Name,
			serverConfig.Db.Prod.Port,
			serverConfig.Db.Prod.SSLMode,
			serverConfig.Db.Prod.Host,
		)
	}
	if serverConfig.Server.App.Mode == "development" {
		connStr = fmt.Sprintf(
			"user=%s password=%s dbname=%s port=%s sslmode=%s host=%s",
			serverConfig.Db.Dev.User,
			serverConfig.Db.Dev.Password,
			serverConfig.Db.Dev.Name,
			serverConfig.Db.Dev.Port,
			serverConfig.Db.Dev.SSLMode,
			serverConfig.Db.Dev.Host,
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
