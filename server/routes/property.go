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
		return sendError(context, "no property information in JSON /property GET")
	}

	property, err = database.GetPropertyPublic(property.Key)
	if err != nil {
		return sendError(context, "no property information in DB /property GET")
	}

	if property.Key == "carusel_id" {
		images, err := database.GetImages(property.Value)
		if err != nil {
			return sendError(context, "can't get images /property GET")
		}
		return context.JSON(http.StatusOK, map[string]interface{}{
			"status":         "success",
			"carusel_images": images,
		})
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
		return sendError(context, "not admin /property GET")
	}

	property := &database.Property{}
	err := context.Bind(property)
	if err != nil {
		return sendError(context, "no property information in JSON /property GET")
	}

	property, err = database.GetPropertyPrivateAndPublic(property.Key)
	if err != nil {
		return sendError(context, "no property information in DB /property GET")
	}

	if property.Key == "carusel_id" {
		images, err := database.GetImages(property.Value)
		if err != nil {
			return sendError(context, "can't get images /property GET")
		}
		return context.JSON(http.StatusOK, map[string]interface{}{
			"status":         "success",
			"carusel_images": images,
		})
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
		"value":  property.Value,
	})
}
