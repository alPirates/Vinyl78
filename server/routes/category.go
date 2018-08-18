package routes

import (
	"net/http"
	"strconv"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func getCategory(context echo.Context) error {

	categories, err := database.GetCategories()
	if err != nil {
		return sendError(context, "categories not returned from db /sidebar GET")
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"result": categories,
		"status": "success",
	})
}

func addCategory(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /addCategory POST")
	}

	category := database.Category{} // Use name, icon
	err := context.Bind(&category)
	if err != nil {
		return sendError(context, "no user information in JSON /addCategory POST")
	}

	if category.Name == "" || category.Icon == "" {
		return sendError(context, "empty params /addCategory POST")
	}

	_, err = category.Create(
		category.Name,
		category.Icon,
	)

	if err != nil {
		return sendError(context, "category not created /addCategory POST")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})

}

func deleteCategory(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /deleteCategory DELETE")
	}

	idParam := context.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return sendError(context, "skip is not uint /deleteCategory DELETE")
	}
	idUint := uint(id)

	category := &database.Category{
		ID: idUint,
	}

	err = category.Delete()
	if err != nil {
		return sendError(context, "category not deleted /deleteCategory DELETE")
	}

	err = database.DeleteByCategory(category.ID)
	if err != nil {
		return sendError(context, "stickers in this category not deleted /deleteCategory DELETE")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})

}
