package routers

import (
	"net/http"
	"sshClient/web/view/components"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/ssh"
)

func CmdRoutes(e *echo.Echo) {
	run := e.Group("/run")
	run.GET("/:cmd", cmdHandler)
}

func cmdHandler(c echo.Context) error {
	if conn == nil {
		component := components.ErrorTempl("No Server Connection")
		return component.Render(c.Request().Context(), c.Response())
	}
	switch cmd := c.Param("cmd"); cmd {
	case "getOs":
		value, err := RunCmd(conn, "uname -s")
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		component := components.GetOperatingSystem(value)
		return component.Render(c.Request().Context(), c.Response())
	}
	return c.JSON(http.StatusOK, "No Data")
}

func RunCmd(connection *ssh.Client, cmd string) (string, error) {
	session, err := connection.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// Run the command
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
