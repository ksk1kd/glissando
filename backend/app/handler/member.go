package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"glissando/store"
)

func GetMemberList(c echo.Context) error {

	members, err := store.GetMembers()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, members)
}
