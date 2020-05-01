package database

import (
	"database/sql"
	"fmt"
)

// Config defines the mysql credentials
type Config struct {
	User     string
	Password string
	Database string
}

// DB whatever
type DB struct {
	Conn *sql.DB
}

// Connect to the database
func Connect(config *Config) (*DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", config.User, config.Password, config.Database)) // "user:password@/dbname"
	return &DB{Conn: db}, err
}
