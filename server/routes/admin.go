package routes

import (
	"net/http"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func getAdminInfo(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /admin GET")
	}

	property, err := database.GetPropertyPrivateAndPublic("carusel_id")
	if err != nil {
		return sendError(context, "no carusel id /admin GET")
	}

	caruselImages, err := database.GetImages(property.Value)
	if err != nil {
		return sendError(context, "cant't get carusel images /admin GET")
	}

	applications, err := database.GetAllApplications()
	if err != nil {
		return sendError(context, "cant't get applications /admin GET")
	}

	categories, err := database.GetCategories()
	if err != nil {
		return sendError(context, "cant't get categories /admin GET")
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"status":         "success",
		"carusel_images": caruselImages,
		"applications":   applications,
		"categories":     categories,
	})
}
