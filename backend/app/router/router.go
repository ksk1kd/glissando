package router

import (
	"github.com/labstack/echo/v4"

	"glissando/handler"
)

func SetRouting(e *echo.Echo) {
	version := "v1"
	prefix := "/api/" + version
	g := e.Group(prefix)
	g.GET("/items", handler.GetItemList)
}
