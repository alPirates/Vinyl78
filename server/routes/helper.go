package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Port - number of a port
const (
	Port = 3000
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
	api.Static("/media", "../media")

	api.GET("/sidebar", getCategory)

	authorization := api.Group("/app")
	authorization.Use(middleware.JWTWithConfig(
		middleware.JWTConfig{
			Claims:     &jwtUserClaims{},
			ContextKey: "token",
			SigningKey: []byte(SignedString),
		}))

	authorization.POST("/unauthorization", unauthorization) // user
	authorization.PUT("/user", setUser)                     // user
	authorization.DELETE("/user", deleteUser)               // user !!!!USE TOKEN TO GET ID!!!!
	authorization.POST("/update", updateToken)              // user

	authorization.PUT("/application", setApplication)           // admin
	authorization.GET("/application", getApplication)           // admin
	authorization.DELETE("/application/:id", deleteApplication) // admin
	authorization.PUT("/token", setToken)                       // admin
	authorization.PUT("/sticker", setSticker)                   // admin
	authorization.DELETE("/sticker/:id", deleteSticker)         // admin
	authorization.POST("/sticker", addSticker)                  // admin
	authorization.DELETE("/category/:id", deleteCategory)       // admin
	authorization.POST("/category", addCategory)                // admin
	authorization.POST("/image", addImage)                      // admin
	authorization.DELETE("/image/:id", deleteImage)             // admin

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
