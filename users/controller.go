package users

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/renato-macedo/issuetracker_go/roles"
	"github.com/renato-macedo/issuetracker_go/utils"
)

// Controller for requests
type Controller struct {
	Service *Service
}

// NewController returns a new Controller
func NewController(service *Service) *Controller {
	return &Controller{
		Service: service,
	}
}

// GetUsers handle a Get to /users
func (controller *Controller) GetUsers(c echo.Context) error {

	role := &roles.Role{ID: 1, Title: "admin", Description: "administrator"}
	user := &User{
		ID:       1,
		Name:     "Renato",
		Email:    "renato@gmail.com",
		Password: "renato123",
		Role:     role,
	}

	return c.JSON(http.StatusOK, user)
}

// GetUser return a single user json response
func (controller *Controller) GetUser(c echo.Context) error {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {

		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid id"))
	}

	user, err := controller.Service.FindUserByID(id)
	if err != nil {
		c.Logger().Printf("%v", err)
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid id"))
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser handles a put on /user
func (controller *Controller) UpdateUser(c echo.Context) error {
	id := c.Param("id")

	return c.String(http.StatusOK, "updated "+id)
}

// DeleteUser handles a delete on /user
func (controller *Controller) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "deleted "+id)
}

// Login controller
func (controller *Controller) Login(c echo.Context) error {
	credentials := &LoginDTO{}
	err := c.Bind(credentials)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("Invalid data"))
	}
	err = controller.Service.CheckCredentials(c.Logger(), credentials)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest(err.Error()))
	}

	return c.JSON(http.StatusOK, struct{ Message string }{Message: "yay!!!"})
}

// Register controller
func (controller *Controller) Register(c echo.Context) error {
	data := &RegisterDTO{}

	err := c.Bind(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("Invalid data"))
	}

	if data.Password != data.ConfirmPassword {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("Passwords must be equal"))
	}

	_, err = controller.Service.FindUserByEmail(data.Email)

	if err == nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("User already exists"))
	}

	err = controller.Service.Register(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewServerError("It was not possible to "))
	}

	return c.JSON(http.StatusCreated, utils.NewOkReponse("user registered"))
}
