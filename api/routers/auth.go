package routers

import (
	"log"
	"sshClient/types"
	"sshClient/web/view/components"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/ssh"
)

var login string = "off"
var data types.LoginConfig
var conn *ssh.Client

func AuthRoutes(e *echo.Echo) {
	// Auth routes
	auth := e.Group("/auth")
	auth.POST("/login", loginHandler)
	auth.GET("/logout", logoutHandler)

}

func loginHandler(c echo.Context) error {
	login = "on"
	if err := c.Bind(&data); err != nil {
		return err
	}
	data.Addr = data.Addr + ":22"
	conn = Connect(data.User, data.Password, data.Addr)
	return c.Redirect(302, "/")
}

func logoutHandler(c echo.Context) error {
	login = "off"
	ConnectionError(c, conn)
	CloseConnection(conn)
	return c.Redirect(302, "/")
}

func Connect(user, password, addr string) *ssh.Client {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	connection, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal(err)
	}

	return connection
}

func CloseConnection(connection *ssh.Client) {
	connection.Close()
}

func ConnectionError(c echo.Context, connection *ssh.Client) error {
	if connection == nil {
		component := components.ErrorTempl("No Server Connection")
		return component.Render(c.Request().Context(), c.Response())
	}
	return nil
}
