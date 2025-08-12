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

func Connect(serverConfig config.Configs) (*sql.DB, error) {
	var connStr string
	if serverConfig.Server.App.Mode == "prod" {
		connStr = fmt.Sprintf(
			"user=%s password=%s dbname=%s port=%s sslmode=%s host=%s",
			serverConfig.Db.Prod.User,
			serverConfig.Db.Prod.Password,
			serverConfig.Db.Prod.Name,
			serverConfig.Db.Prod.Port,
			serverConfig.Db.Prod.SSLMode,
			serverConfig.Db.Prod.Host,
		)
	} else {
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

	var lastErr error
	for i := range 5 {
		fmt.Println("Database connection attempt", i+1)

		conn, err := sql.Open("postgres", connStr)
		if err != nil {
			lastErr = err
			log.Printf("Failed to open connection: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = conn.Ping()
		if err == nil {
			return conn, nil
		}
		conn.Close()

		if strings.Contains(err.Error(), "database system is starting up") {
			log.Println("Database is starting up, retrying...")
		} else {
			log.Printf("Ping failed: %v", err)
		}

		lastErr = err
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("connection failed after retries: %w", lastErr)
}
