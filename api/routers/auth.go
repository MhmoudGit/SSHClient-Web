package routers

import (
	"sshClient/web/view"

	"github.com/labstack/echo/v4"
)

var login string = "off"

func MainRoutes(e *echo.Echo) {
	e.GET("/", mainHandler)

	// Auth routes
	auth := e.Group("/auth")
	auth.GET("/login", loginHandler)
	auth.GET("/logout", logoutHandler)

}

func mainHandler(c echo.Context) error {
	home := view.Home()
	component := view.Base(home, login)
	return component.Render(c.Request().Context(), c.Response())
}
func loginHandler(c echo.Context) error {
	login = "on"
	return c.Redirect(302, "/")
}

func logoutHandler(c echo.Context) error {
	login = "off"
	return c.Redirect(302, "/")
}
