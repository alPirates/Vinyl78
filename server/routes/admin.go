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
		return sendError(context, "not admin", "вы не администратор")
	}

	property, err := database.GetPropertyPrivateAndPublic("carousel_id")
	if err != nil {
		return sendError(context, "no carousel_id property", "не удалось получить информацию")
	}

	caruselImages, err := database.GetImages(property.Value)
	if err != nil {
		return sendError(context, "cant't get carusel images", "не удалось получить информацию")
	}

	applications, err := database.GetAllApplications()
	if err != nil {
		return sendError(context, "cant't get applications", "не удалось получить информацию")
	}

	categories, err := database.GetCategories()
	if err != nil {
		return sendError(context, "cant't get categories", "не удалось получить информацию")
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"status":         "success",
		"carusel_images": caruselImages,
		"applications":   applications,
		"categories":     categories,
		"message":        "информация получена",
	})
}
