package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Port - number of a port
// SignedString - signed string
const (
	Port         = 8080
	SignedString = "secret"
)

// User structure
type User struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

// SetRoutes function
// Set all routes
func SetRoutes(server *echo.Echo) {

	server.GET("/", getHTML)

	// appGroup := server.Group("/app")
	// appGroup.GET("", nil)

	server.POST("/authorization", authorization)
	server.POST("/registration", registration)
	server.POST("/unauthorization", unauthorization)

	userGroup := server.Group("/user", middleware.JWT([]byte(SignedString)))
	userGroup.POST("/", setUser)
	userGroup.PUT("/", addUser)
	userGroup.DELETE("/", deleteUser)

	adminGroup := server.Group("/user", nil)
	adminGroup.POST("/admin/token", setToken)

	err := server.Start(":" + fmt.Sprint(Port))
	if err != nil {
		server.Logger.Fatal(err)
	}

}

func registration(context echo.Context) error {

	userInfo := User{}
	err := context.Bind(userInfo)
	if err != nil {
		return sendError(context, "no user information in JSON /registration")
	}

	u := database.GetUserFromEmail(userInfo.Email)
	if u != nil {
		return sendError(context, "email already exist /registration")
	}
	user := database.AddUser(userInfo.Email, userInfo.Password, false)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = user.GetID()
	claims["email"] = user.GetEmail()
	claims["isAdmin"] = user.GetRole()
	claims["expirationDate"] = time.Now().Add(time.Hour * 72).Unix()

	endToken, err := token.SignedString([]byte(SignedString))
	if err != nil {
		return sendError(context, "token not encoded /registration")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"token":  endToken,
		"status": "success",
	})
}

func authorization(context echo.Context) error {

	userInfo := User{}
	err := context.Bind(&userInfo)
	if err != nil {
		return sendError(context, "no user information in JSON /authorization")
	}

	user := database.GetUser(userInfo.Email, userInfo.Password)
	if user == nil {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = user.GetID()
	claims["email"] = user.GetEmail()
	claims["isAdmin"] = user.GetRole()
	claims["expirationDate"] = time.Now().Add(time.Hour * 72).Unix()

	endToken, err := token.SignedString([]byte(SignedString))
	if err != nil {
		return sendError(context, "token not encoded /authorizationa")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"token":  endToken,
		"status": "success",
	})
}

func unauthorization(context echo.Context) error {
	return nil
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

func sendError(context echo.Context, errorName string) error {
	return context.JSON(http.StatusOK, map[string]string{
		"status": "failure",
		"error":  errorName,
	})
}
