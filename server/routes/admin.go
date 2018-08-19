package routes

import (
	"net/http"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func setToken(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /setToken POST")
	}

	property := &database.Property{}
	err := context.Bind(property)
	if err != nil {
		return sendError(context, "no user information in JSON /setToken POST")
	}

	property.Key = "telegram_token"
	err = property.Update()
	if err != nil {
		return sendError(context, "not update token /setToken POST")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}
