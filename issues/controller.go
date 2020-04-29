package issues

import (
	"net/http"

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
	issues, err := controller.Service.FindIssues()
	if err != nil {
		c.Logger().Printf("\n %v", err)
		return c.JSON(http.StatusInternalServerError, utils.NewServerError(err.Error()))
	}
	return c.JSON(http.StatusOK, issues)
}

// CreateIssue controller
func (controller *Controller) CreateIssue(c echo.Context) error {
	return c.String(http.StatusOK, "create issue")
}

// UpdateIssue controller
func (controller *Controller) UpdateIssue(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "update issues "+id)
}

// DeleteIssue controller
func (controller *Controller) DeleteIssue(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "issues"+id)
}
