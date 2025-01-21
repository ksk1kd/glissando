package router

import (
	"github.com/labstack/echo/v4"

	"glissando/handler"
)

func SetRouting(e *echo.Echo) {
	version := "v1"
	prefix := "/api/" + version
	g := e.Group(prefix)
	g.GET("/members", handler.GetMemberList)
	g.GET("/projects", handler.GetProjectList)
	g.GET("/resources", handler.GetResourceList)
	g.PUT("/resources", handler.UpdateResource)
}
