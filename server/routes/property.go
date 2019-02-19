package routes

import (
	"fmt"
	"net/http"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func incrementSeoLink(context echo.Context) error {
	property := &database.Property{}
	err := context.Bind(property)

	if err != nil {
		return sendError(context, "no property information in JSON", "")
	}
	fmt.Println("console is ", property.Key)

	fl := database.IncrementSeoLink(property.Key)
	if !fl {
		return sendError(context, "Cant increment", "")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}

func setProperty(context echo.Context) error {
	property := &database.Property{}
	err := context.Bind(property)

	fmt.Println("prop", property)

	if err != nil {
		return sendError(context, "no property information in JSON", "")
	}
	fmt.Println("finding by key", property.Key)
	pr, err := database.GetPropertyPrivateAndPublic(property.Key)

	if err != nil {
		return sendError(context, "Cant find property", "")
	}

	pr.Value = property.Value
	err = pr.Save()

	if err != nil {
		return sendError(context, "Cant save property", "")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}

func getProperty(context echo.Context) error {

	property := &database.Property{}
	err := context.Bind(property)
	if err != nil {
		return sendError(context, "no property information in JSON", "")
	}

	property, err = database.GetPropertyPublic(property.Key)
	if err != nil {
		return sendError(context, "no property information", "")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
		"value":  property.Value,
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
		return sendError(context, "no property information in JSON", "")
	}

	property, err = database.GetPropertyPrivateAndPublic(property.Key)
	if err != nil {
		return sendError(context, "no property information", "")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
		"value":  property.Value,
	})
}
