package routes

import (
	"net/http"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func getProperty(context echo.Context) error {

	property := &database.Property{}
	err := context.Bind(property)
	if err != nil {
		return sendError(context, "no property information in JSON", "не удалось получить свойство")
	}

	property, err = database.GetPropertyPublic(property.Key)
	if err != nil {
		return sendError(context, "no property information", "не удалось получить свойство")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"value":   property.Value,
		"message": "свойство получено",
	})

}

func getPrivateProperty(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin", "вы не администратор")
	}

	property := &database.Property{}
	err := context.Bind(property)
	if err != nil {
		return sendError(context, "no property information in JSON", "не удалось получить свойство")
	}

	property, err = database.GetPropertyPrivateAndPublic(property.Key)
	if err != nil {
		return sendError(context, "no property information", "не удалось получить свойство")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"value":   property.Value,
		"message": "свойство получено",
	})
}
