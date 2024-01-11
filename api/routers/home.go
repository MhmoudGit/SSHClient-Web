package routers

import (
	"sshClient/web/view"

	"sshClient/web/view/pages"

	"github.com/labstack/echo/v4"
)

func HomeRoutes(e *echo.Echo) {
	e.GET("/", mainHandler)
}

func mainHandler(c echo.Context) error {
	home := pages.Home(data)
	component := view.Base(home, login)
	return component.Render(c.Request().Context(), c.Response())
}
