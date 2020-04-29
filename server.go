package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/renato-macedo/issuetracker_go/database"
	"github.com/renato-macedo/issuetracker_go/issues"
	"github.com/renato-macedo/issuetracker_go/roles"
	"github.com/renato-macedo/issuetracker_go/tags"
	"github.com/renato-macedo/issuetracker_go/users"
)

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

func startServer(db *database.DB) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})

	// User
	userService := users.NewService(db)
	userController := users.NewController(userService)
	e.GET("/users", userController.GetUsers)
	e.GET("/users/:id", userController.GetUser)
	e.PUT("/users/:id", userController.UpdateUser)
	e.DELETE("/users/:id", userController.DeleteUser)

	// Auth
	e.POST("/auth/login", userController.Login)
	e.POST("/auth/register", userController.Register)

	// Roles
	roleController := &roles.RoleController{}
	e.GET("/roles", roleController.GetRoles)

	// Issues
	issuesController := issues.NewController(db)
	e.GET("/issues", issuesController.GetIssues)
	e.POST("/issues", issuesController.CreateIssue)
	e.PUT("/issues/:id", issuesController.UpdateIssue)
	e.DELETE("/issues/:id", issuesController.DeleteIssue)
	// Tags
	tagController := &tags.TagController{}
	e.GET("/tags", tagController.GetTags)
	e.GET("/tags/", tagController.GetTags)
	e.Logger.Fatal(e.Start(":5000"))
}
