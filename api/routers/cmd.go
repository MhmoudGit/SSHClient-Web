package routers

import (
	"net/http"
	"sshClient/types"
	"sshClient/web/view/components"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/ssh"
)

func CmdRoutes(e *echo.Echo) {
	run := e.Group("/run")
	run.GET("/domains", domainHandler)
	run.GET("/:cmd", cmdHandler)
}

func extractDomains(input string) []string {
	// Split the input string based on spaces
	domainList := strings.Fields(input)
	return domainList
}

func domainHandler(c echo.Context) error {
	ConnectionError(c, conn)
	value, err := RunCmd(conn, types.DomainsCmd)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	domains := extractDomains(value)
	component := components.DomainsData(domains)
	return component.Render(c.Request().Context(), c.Response())
}

func cmdHandler(c echo.Context) error {
	ConnectionError(c, conn)
	switch cmd := c.Param("cmd"); cmd {
	case "getOs":
		value, err := RunCmd(conn, types.OsCmd)
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		component := components.GetValue(value)
		return component.Render(c.Request().Context(), c.Response())
	case "getHost":
		value, err := RunCmd(conn, types.HostCmd)
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		component := components.GetValue(value)
		return component.Render(c.Request().Context(), c.Response())
	case "getIp":
		value, err := RunCmd(conn, types.IpCmd)
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		component := components.GetValue(value)
		return component.Render(c.Request().Context(), c.Response())
	case "getSpecs":
		value, err := RunCmd(conn, types.RamCmd)
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		value2, err := RunCmd(conn, types.CpuCmd)
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
