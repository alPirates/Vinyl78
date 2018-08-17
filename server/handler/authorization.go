package handler

import (
	"net/http"
	"time"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func registration(context echo.Context) error {

	userInfo := database.User{}
	err := context.Bind(&userInfo)
	if err != nil {
		return sendError(context, "no user information in JSON /registration")
	}

	if database.CheckEmailAndPassword(userInfo.Email, userInfo.Password) {
		return sendError(context, "email already existed /registration")
	}
	user := (&database.User{}).Create(userInfo.Email, userInfo.Password)

	claims := &jwtUserClaims{
		ID:   user.Model.ID,
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

	claims := &jwtUserClaims{
		ID:   user.Model.ID,
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

	return context.JSON(http.StatusOK, map[string]string{
		"token":  endToken,
		"status": "success",
	})
}

func unauthorization(context echo.Context) error {
	return nil
}
