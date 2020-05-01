package issues

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/renato-macedo/issuetracker_go/database"
	"github.com/renato-macedo/issuetracker_go/utils"
)

// Controller for handle requests
type Controller struct {
	Service *Service
}

// NewController creates a new Controller
func NewController(db *database.DB) *Controller {
	return &Controller{
		Service: &Service{
			DB: db,
		},
	}
}

// GetIssues controller
func (controller *Controller) GetIssues(c echo.Context) error {
	param := c.QueryParam("closed")
	includeClosed := false
	var err error
	if param != "" {
		includeClosed, err = strconv.ParseBool(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.NewBadRequest("Invalid query param"))
		}
	}

	issues, err := controller.Service.FindIssues(includeClosed)
	if err != nil {
		c.Logger().Printf("\n %v", err)
		return c.JSON(http.StatusInternalServerError, utils.NewServerError(err.Error()))
	}

	return c.JSON(http.StatusOK, issues)
}

// CreateIssue controller
func (controller *Controller) CreateIssue(c echo.Context) error {
	data := &IssueDTO{}

	err := c.Bind(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("Invalid body"))
	}
	err = c.Validate(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("Invalid body"))
	}

	return c.String(http.StatusOK, "kkkk")
}

// UpdateIssue controller
func (controller *Controller) UpdateIssue(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusNotImplemented, "update issues "+id)
}

// DeleteIssue controller
func (controller *Controller) DeleteIssue(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusNotImplemented, "issues"+id)
}
