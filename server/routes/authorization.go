package routes

import (
	"net/http"
	"time"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func registration(context echo.Context) error {

	userInfo := database.User{}
	err := context.Bind(&userInfo) // email and password
	if err != nil {
		return sendError(context, "no user information in JSON /registration")
	}

	if userInfo.Email == "" || userInfo.Password == "" {
		return sendError(context, "empty params /registration")
	}

	if database.CheckEmailAndPassword(userInfo.Email, userInfo.Password) {
		return sendError(context, "email already existed /registration")
	}
	user, err := (&database.User{}).Create(userInfo.Email, userInfo.Password)

	if err != nil {
		return sendError(context, "user not created /registration")
	}

	claims := &jwtUserClaims{
		UUID: user.ID,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	endToken, err := token.SignedString([]byte(SignedString))
	if err != nil {
		return sendError(context, "token not encoded /registration")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"token":  endToken,
		"role":   "user",
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
	user, err := database.GetUser(userInfo.Email, userInfo.Password)

	if err != nil {
		return sendError(context, "can't get user from db /authorization")
	}

	claims := &jwtUserClaims{
		UUID: user.ID,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	endToken, err := token.SignedString([]byte(SignedString))
	if err != nil {
		return sendError(context, "token not encoded /authorizationa")
	}

	var roleStr string
	if user.Role == 1 {
		roleStr = "admin"
	} else {
		roleStr = "user"
	}

	return context.JSON(http.StatusOK, map[string]string{
		"token":  endToken,
		"role":   roleStr,
		"status": "success",
	})
}

func unauthorization(context echo.Context) error {
	return nil
}

func updateToken(context echo.Context) error {
	return nil
}
