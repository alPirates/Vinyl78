package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Port - number of a port use 80 port on proxy
const (
	// Port = 3000
	Port = 80
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

	api.GET("/carousel", getCarousel)
	api.GET("/sticker", getSticker)
	api.GET("/sticker/:id", getStickerByID)
	api.GET("/media/:name", getFileImage)

	api.GET("/sidebar", getCategory)
	api.GET("/main_sidebar", getCategories)
	api.GET("/image", getImage)
	api.GET("/property", getProperty)
	api.GET("/category/:id", getCategoryByID)
	api.PUT("/inc_seo", incrementSeoLink)

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

	authorization.PUT("/carousel", setСarousel)
	authorization.PUT("/slides", updateNumber) // update slides position number
	authorization.DELETE("/carousel/:id", deleteCarousel)
	authorization.POST("/carousel", addCarousel)

	authorization.PUT("/application", setApplication)                    // admin
	authorization.GET("/application", getApplication)                    // admin
	authorization.DELETE("/application/:id", deleteApplication)          // admin
	authorization.GET("/property", getPrivateProperty)                   // admin
	authorization.PUT("/property", setProperty)                          // admin
	authorization.GET("/admin", getAdminInfo)                            // admin
	authorization.PUT("/sticker", setSticker)                            // admin
	authorization.DELETE("/sticker/:id", deleteSticker)                  // admin
	authorization.POST("/sticker", addSticker)                           // admin
	authorization.PUT("/updateStickersPosition", updateStickersPosition) // admin
	authorization.DELETE("/category/:id", deleteCategory)                // admin
	authorization.PUT("/category", setCategory)                          // admin
	authorization.PUT("/categories", setCategories)                      // admin
	authorization.POST("/category", addCategory)                         // admin
	authorization.PUT("/categories_number", updateCategoryNumber)
	authorization.POST("/image", addImage)          // admin
	authorization.PUT("/image", setImage)           // admin
	authorization.DELETE("/image/:id", deleteImage) // admin

	err := server.Start(":" + fmt.Sprint(Port))
	if err != nil {
		server.Logger.Fatal(err)
	}

}

func sendError(context echo.Context, errorName, message string) error {
	fmt.Println(time.Now().Format(time.UnixDate) + " : " + context.Path() + " : " + errorName)
	return context.JSON(http.StatusOK, map[string]string{
		"status":  "failure",
		"error":   errorName,
		"message": message,
	})
}
