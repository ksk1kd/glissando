package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"glissando/store"
)

func GetResourceList(c echo.Context) error {

	member := c.QueryParam("member")
	project := c.QueryParam("project")
	month := c.QueryParam("month")

	resources, err := store.GetResources(member, project, month)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, resources)
}
