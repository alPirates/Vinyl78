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

// jwtCustomClaims are custom claims extending default ones.
type jwtUserClaims struct {
	jwt.StandardClaims
	ID       string `json:"ID"`
	Email    bool   `json:"email"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}

// SetRoutes function
// Set all routes
func SetRoutes(server *echo.Echo) {

	server.File("/", "../web/dist/index.html")

	// appGroup := server.Group("/app")
	// appGroup.GET("", nil)

	server.POST("/authorization", authorization)
	server.POST("/registration", registration)
	server.POST("/unauthorization", unauthorization)

	userGroup := server.Group("/user", middleware.JWT(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwtUserClaims{},
		ContextKey: "token",
		SigningKey: []byte(SignedString),
	})))
	userGroup.POST("", setUser)
	userGroup.DELETE("", deleteUser)

	// adminGroup := server.Group("/user", nil)
	// adminGroup.POST("/admin/token", setToken)

	err := server.Start(":" + fmt.Sprint(Port))
	if err != nil {
		server.Logger.Fatal(err)
	}

}

func registration(context echo.Context) error {

	userInfo := database.User{}
	err := context.Bind(&userInfo)
	if err != nil {
		return sendError(context, "no user information in JSON /registration")
	}

	if database.CheckEmailAndPassword(userInfo.Email, userInfo.Password) {
		return sendError(context, "email already existed /registration")
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

	userInfo := database.User{}
	err := context.Bind(&userInfo)
	if err != nil {
		return sendError(context, "no user information in JSON /authorization")
	}

	if !database.CheckEmailAndPassword(userInfo.Email, userInfo.Password) {
		return sendError(context, "no user information in DB /authorization")
	}
	user := database.GetUser(userInfo.Email, userInfo.Password)

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

func setUser(context echo.Context) error {
	tokenFromWeb := context.Get("token").(*jwt.Token)

	claimsFromWeb := tokenFromWeb.Claims.(jwt.MapClaims)
	ID := claimsFromWeb["ID"].(uint)
	email := claimsFromWeb["email"].(string)
	password := claimsFromWeb["password"].(string)
	role := claimsFromWeb["isAdmin"].(uint)

	if !database.CheckID(ID) {
		return sendError(context, "no user information in DB /user/ POST")
	}

	user := database.GetUserByID(ID)
	if role == 1 {
		user.Set(email, password, true)
	} else {
		user.Set(email, password, false)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = user.GetID()
	claims["email"] = user.GetEmail()
	claims["role"] = user.GetRole()
	claims["expirationDate"] = time.Now().Add(time.Hour * 72).Unix()

	endToken, err := token.SignedString([]byte(SignedString))
	if err != nil {
		return sendError(context, "token not encoded /user/ POST")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"token":  endToken,
		"status": "success",
	})
}

func deleteUser(context echo.Context) error {
	tokenFromWeb := context.Get("token").(*jwt.Token)

	claimsFromWeb := tokenFromWeb.Claims.(jwt.MapClaims)
	ID := claimsFromWeb["ID"].(uint)

	if !database.CheckID(ID) {
		return sendError(context, "no user information in DB /user/ DELETE")
	}

	user := database.GetUserByID(ID)
	user.Delete()

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}

func getAdmin(context echo.Context) error {
	return context.String(http.StatusOK, "This is admin information")
}

func setToken(context echo.Context) error {
	return context.String(http.StatusOK, "This can set TelegToken")
}

func sendError(context echo.Context, errorName string) error {
	return context.JSON(http.StatusOK, map[string]string{
		"status":  "failure",
		"message": errorName,
	})
}
