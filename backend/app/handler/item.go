package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"glissando/store"
)

func GetItemList(c echo.Context) error {

	q := c.QueryParam("q")
	category := c.QueryParam("category")
	if category == "0" {
		category = ""
	}

	items, err := store.SearchItems(q, category)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, items)
}
