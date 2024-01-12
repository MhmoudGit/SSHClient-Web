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
	ConnectionError(c, conn)
	switch cmd := c.Param("cmd"); cmd {
	case "getOs":
		value, err := RunCmd(conn, "lsb_release -a 2>/dev/null | grep 'Description' | awk -F ':\t' '{print $2}'")
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		component := components.GetValue(value)
		return component.Render(c.Request().Context(), c.Response())
	case "getHost":
		value, err := RunCmd(conn, "hostname")
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		component := components.GetValue(value)
		return component.Render(c.Request().Context(), c.Response())
	case "getIp":
		value, err := RunCmd(conn, "hostname -I | awk '{print $1}'")
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		component := components.GetValue(value)
		return component.Render(c.Request().Context(), c.Response())
	case "getRam":
		value, err := RunCmd(conn, "sudo dmidecode -t 17 | grep Size")
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		value2, err := RunCmd(conn, "nproc")
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		value3 := value2 + "Cores," + value + "Memory"
		component := components.GetValue(value3)
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
