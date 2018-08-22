package routes

import (
	"net/http"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func setUser(context echo.Context) error {

	user := &database.User{}
	err := context.Bind(user)
	if err != nil {
		return sendError(context, "no user information in JSON /user PUT")
	}

	err = user.Update()
	if err != nil {
		return sendError(context, "can't update user /user PUT")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
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
		return sendError(context, "application not deleted /application DELETE")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}
