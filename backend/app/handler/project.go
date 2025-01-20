package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"glissando/store"
)

func GetProjectList(c echo.Context) error {

	projects, err := store.GetProjects()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, projects)
}
