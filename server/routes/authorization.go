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
		return sendError(context, "no user information in JSON", "не удалось зарегистровать пользователя")
	}

	if userInfo.Email == "" || userInfo.Password == "" {
		return sendError(context, "empty params", "не удалось зарегистровать пользователя")
	}

	if database.CheckEmailAndPassword(userInfo.Email, userInfo.Password) {
		return sendError(context, "email already exist", "E-mail занят")
	}
	user, err := (&database.User{}).Create(userInfo.Email, userInfo.Password)

	if err != nil {
		return sendError(context, "user not created", "не удалось зарегистровать пользователя")
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
		return sendError(context, "token not encoded", "не удалось зарегистровать пользователя")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"token":   endToken,
		"role":    "user",
		"status":  "success",
		"message": "пользователь зарегистрирован",
	})
}

func authorization(context echo.Context) error {

	userInfo := database.User{}
	err := context.Bind(&userInfo)
	if err != nil {
		return sendError(context, "no user information in JSON", "не удалось авторизировать пользователя")
	}

	if !database.CheckEmailAndPassword(userInfo.Email, userInfo.Password) {
		return sendError(context, "no this user", "не удалось авторизировать пользователя")
	}
	user, err := database.GetUser(userInfo.Email, userInfo.Password)

	if err != nil {
		return sendError(context, "can't get user", "не удалось авторизировать пользователя")
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
		return sendError(context, "token not encoded", "не удалось авторизировать пользователя")
	}

	var roleStr string
	if user.Role == 1 {
		roleStr = "admin"
	} else {
		roleStr = "user"
	}

	return context.JSON(http.StatusOK, map[string]string{
		"token":   endToken,
		"role":    roleStr,
		"status":  "success",
		"message": "пользователь авторизован",
	})
}

func unauthorization(context echo.Context) error {
	return nil
}

func updateToken(context echo.Context) error {
	return nil
}
