package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Dattito/HMN_backend_api/app/utils/env_utils"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
)

// PostgreSQLConnection func for connection to PostgreSQL database.
func PostgreSQLConnection() (*sqlx.DB, error) {
	// Define database connection settings.
	maxConn, _ := strconv.Atoi(env_utils.GetEnv("DB_MAX_CONNECTIONS", "0"))
	maxIdleConn, _ := strconv.Atoi(env_utils.GetEnv("DB_MAX_IDLE_CONNECTIONS", "2"))
	maxLifetimeConn, _ := strconv.Atoi(env_utils.GetEnv("DB_MAX_LIFETIME_CONNECTIONS", "0"))

	// Define database connection for PostgreSQL.
	db, err := sqlx.Connect("pgx", os.Getenv("DB_SERVER_URL"))
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Set database connection settings.
	db.SetMaxOpenConns(maxConn)                           // the default is 0 (unlimited)
	db.SetMaxIdleConns(maxIdleConn)                       // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) // 0, connections are reused forever

	return db, nil
}
