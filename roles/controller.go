package roles

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// RoleController for requests
type RoleController struct{}

// GetRoles handles get on /roles
func (controller *RoleController) GetRoles(c echo.Context) error {
	roles := make([]*Role, 3)
	roles[0] = &Role{ID: 1, Title: "admin", Description: "can assign issues to users"}
	roles[1] = &Role{ID: 2, Title: "employee", Description: "takes the issues"}
	roles[2] = &Role{ID: 3, Title: "guest", Description: "the guest can only watch"}
	return c.JSON(http.StatusOK, roles)
}
