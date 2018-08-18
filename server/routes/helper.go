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
	server.Static("/static", "../web/dist/static")

	api := server.Group("/api")

	api.POST("/authorization", authorization)
	api.POST("/registration", registration)

	api.POST("/application", addApplication)

	api.GET("/sticker", getSticker)

	authorization := api.Group("/app")
	authorization.Use(middleware.JWTWithConfig(
		middleware.JWTConfig{
			Claims:     &jwtUserClaims{},
			ContextKey: "token",
			SigningKey: []byte(SignedString),
		}))

	authorization.POST("/unauthorization", unauthorization) // user
	authorization.PUT("/user", setUser)                     // user
	authorization.DELETE("/user", deleteUser)               // user
	authorization.POST("/update", updateToken)              // user

	authorization.PUT("/application", setApplication)       // admin
	authorization.GET("/application", getApplication)       // admin
	authorization.DELETE("/application", deleteApplication) // admin
	authorization.PUT("/token", setToken)                   // admin
	authorization.PUT("/sticker", setSticker)               // admin
	authorization.DELETE("/sticker", deleteSticker)         // admin
	authorization.POST("/sticker", addSticker)              // admin

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
