package routes

import (
	"net/http"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func setUser(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	user := &database.User{}
	err := context.Bind(user)
	if err != nil {
		return sendError(context, "no user information in JSON", "не удалось изменить пользователя")
	}

	user.ID = claims.UUID
	user.Role = claims.Role

	err = user.UpdateNotAll()
	if err != nil {
		return sendError(context, "can't update user", "не удалось изменить пользователя")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "пользователь изменен",
	})
}

func deleteUser(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	user := &database.User{
		ID: claims.UUID,
	}

	err := user.Delete()

	if err != nil {
		return sendError(context, "can't delete user", "не удалось удалить пользователя")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "пользователь удален",
	})
}
