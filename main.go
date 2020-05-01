package main

import (
	"log"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/renato-macedo/issuetracker_go/database"
)

// CustomValidator for the framework
type CustomValidator struct {
	validator *validator.Validate
}

// Validate function
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {

	config := &database.Config{
		User:     "root",
		Password: "",
		Database: "issuetracker",
	}

	db, err := database.Connect(config)

	if err != nil {
		log.Fatalf("Could not conect to database %v", err)
	}
	startServer(db)
}
