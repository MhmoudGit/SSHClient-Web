package api

import (
	"sshClient/api/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func App() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routers.MainRoutes(e)

	e.Logger.Fatal(e.Start("localhost:3000"))
}
