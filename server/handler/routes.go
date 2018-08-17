package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Port - number of a port
const (
	Port = 8080
)

// SetRoutes function
// Set all routes
func SetRoutes(server *echo.Echo) {

	server.File("/", "../web/dist/index.html")
	server.Static("/static", "")

	// appGroup := server.Group("/app")
	// appGroup.GET("", nil)

	server.POST("/authorization", authorization)
	server.POST("/registration", registration)

	server.POST("/application", addApplication)

	authorization := server.Group("/user", middleware.JWT(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwtUserClaims{},
		ContextKey: "token",
		SigningKey: []byte(SignedString),
	})))

	authorization.POST("/unauthorization", unauthorization) // user
	authorization.PUT("", setUser)                          // user
	authorization.DELETE("", deleteUser)                    // user

	authorization.PUT("/application", setApplication)       // admin
	authorization.GET("/application", getApplication)       // admin
	authorization.DELETE("/application", deleteApplication) // admin
	authorization.PUT("/token", setToken)                   // admin

	err := server.Start(":" + fmt.Sprint(Port))
	if err != nil {
		server.Logger.Fatal(err)
	}

}

func sendError(context echo.Context, errorName string) error {
	return context.JSON(http.StatusOK, map[string]string{
		"status":  "failure",
		"message": errorName,
	})
}
