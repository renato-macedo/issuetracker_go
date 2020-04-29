package tags

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// TagController for requests
type TagController struct{}

// GetTags handle GET on /tags
func (controller *TagController) GetTags(c echo.Context) error {
	tag := &Tag{ID: 1, Title: "Important", Description: ""}
	return c.JSON(http.StatusOK, tag)
}
