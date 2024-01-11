package routers

import (
	// "log"
	"sshClient/web/view"
	"sshClient/web/view/pages"

	"github.com/labstack/echo/v4"
	// "golang.org/x/crypto/ssh"
)

var login string = "off"

var data pages.LoginConfig

func MainRoutes(e *echo.Echo) {
	e.GET("/", mainHandler)

	// Auth routes
	auth := e.Group("/auth")
	auth.POST("/login", loginHandler)
	auth.GET("/logout", logoutHandler)

}

func mainHandler(c echo.Context) error {
	home := pages.Home(data)
	component := view.Base(home, login)
	return component.Render(c.Request().Context(), c.Response())
}

func loginHandler(c echo.Context) error {
	login = "on"
	var u pages.LoginConfig
	if err := c.Bind(&u); err != nil {
		return err
	}
	u.Addr = u.Addr + ":22"
	data = u
	return c.Redirect(302, "/")
}

func logoutHandler(c echo.Context) error {
	login = "off"
	return c.Redirect(302, "/")
}

// func Connect(user, password, addr string) *ssh.Client {
// 	config := &ssh.ClientConfig{
// 		User: user,
// 		Auth: []ssh.AuthMethod{
// 			ssh.Password(password),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 	}

// 	connection, err := ssh.Dial("tcp", addr, config)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return connection
// }

// func CloseConnection(connection *ssh.Client) {
// 	connection.Close()
// }
