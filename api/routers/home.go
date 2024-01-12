package routers

import (
	"sshClient/web/view"

	"sshClient/web/view/components"
	"sshClient/web/view/pages"

	"github.com/labstack/echo/v4"
)

func HomeRoutes(e *echo.Echo) {
	e.GET("/", mainHandler)
	e.GET("/close", closeHandler)
}

func mainHandler(c echo.Context) error {
	home := pages.Home(data)
	component := view.Base(home, login)
	return component.Render(c.Request().Context(), c.Response())
}

func closeHandler(c echo.Context) error {
	component := components.Close()
	return component.Render(c.Request().Context(), c.Response())
}
