package main

import (
	"flag"
	"os"

	"github.com/labstack/echo/v4"

	"glissando/db"
	"glissando/router"
	"glissando/setting"
)

func main() {

	seeding := flag.Bool("seeding", false, "bool flag")
	flag.Parse()

	db.Init(*seeding)

	e := echo.New()

	setting.SetSetting(e)

	router.SetRouting(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}
