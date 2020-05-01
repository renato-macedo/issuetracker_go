package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/renato-macedo/issuetracker_go/database"
	"github.com/renato-macedo/issuetracker_go/issues"
	"github.com/renato-macedo/issuetracker_go/roles"
	"github.com/renato-macedo/issuetracker_go/tags"
	"github.com/renato-macedo/issuetracker_go/users"
)

func startServer(db *database.DB) {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}
	e.Pre(middleware.AddTrailingSlash())

	e.Use(middleware.Logger())

	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})

	// User
	userService := users.NewService(db)
	userController := users.NewController(userService)
	e.GET("/users/", userController.GetUsers)
	e.GET("/users/:id", userController.GetUser)
	e.PUT("/users/:id", userController.UpdateUser)
	e.DELETE("/users/:id", userController.DeleteUser)

	// Auth
	e.POST("/auth/login", userController.Login)
	e.POST("/auth/register", userController.Register)

	// Roles
	roleController := &roles.RoleController{}
	e.GET("/roles/", roleController.GetRoles)

	// Issues
	issuesController := issues.NewController(db)

	var a echo.HandlerFunc
	e.GET("/issues/", issuesController.GetIssues)
	e.GET("/issues/", a)
	e.POST("/issues/", func(c echo.Context) error {
		return issuesController.CreateIssue(c)
	})
	e.PUT("/issues/:id", issuesController.UpdateIssue)
	e.DELETE("/issues/:id", issuesController.DeleteIssue)
	// Tags
	tagController := &tags.TagController{}
	e.GET("/tags/", tagController.GetTags)

	go func() {
		if err := e.Start(":5000"); err != nil {
			e.Logger.Info("error on start, shutting down")
		}
	}()

	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt)
	<-exit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
