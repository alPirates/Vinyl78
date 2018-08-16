package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// PORT - number of a port
const (
	PORT = 8080
)

// SetRoutes function
// Set all routes
func SetRoutes(server *echo.Echo) {

	server.GET("/", getHTML)

	server.POST("/authorization", getUser)
	server.POST("/registration", getUser)
	server.POST("/user", setUser)
	server.PUT("/user", addUser)
	server.DELETE("/user", deleteUser)

	server.GET("/admin", getAdmin)
	server.POST("/admin/token", setToken)
	// TODO: admin functions

	err := server.Start(":" + fmt.Sprint(PORT))
	if err != nil {
		server.Logger.Fatal(err)
	}

}

func getHTML(context echo.Context) error {
	return context.String(http.StatusOK, "This is index.html")
}

func getUser(context echo.Context) error {
	return context.String(http.StatusOK, "This is user information")
}

func setUser(context echo.Context) error {
	return context.String(http.StatusOK, "This can update user")
}

func addUser(context echo.Context) error {
	return context.String(http.StatusOK, "This can add user")
}

func deleteUser(context echo.Context) error {
	return context.String(http.StatusOK, "This can delete user")
}

func getAdmin(context echo.Context) error {
	return context.String(http.StatusOK, "This is admin information")
}

func setToken(context echo.Context) error {
	return context.String(http.StatusOK, "This can set TelegToken")
}
