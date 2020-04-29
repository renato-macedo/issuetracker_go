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

// // CreateIssue inserts into database a issue with status "OPEN" and fill createdAt with the Now() function
// func (db *DB) CreateIssue(issue *issues.IssueDTO) error {
// 	q := `
// 	INSERT INTO ISSUES (title, description, author, status, createdAt) values (?, ?, ?, "OPEN", Now())
// 	`
// 	_, err := db.Conn.Exec(q, issue.Title, issue.Description, issue.AuthorID)
// 	return err
// }
