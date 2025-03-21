package db

import (
	"database/sql"
	"fmt"
	"log"

	"analytics_project/config"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL() (*sql.DB, error) {
	config.LoadEnv()
	dsn := config.GetMySQLDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Connected to MySQL successfully!")
	return db, nil
}
