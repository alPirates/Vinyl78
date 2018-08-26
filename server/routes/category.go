package routes

import (
	"fmt"
	"net/http"

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
		"status": "success",
		"result": categories,
	})
}

func getCategoryByID(context echo.Context) error {
	uuidParam := context.Param("id")
	fmt.Println("id is ", uuidParam)
	result, err := database.GetCategoryByID(uuidParam)
	if err != nil {
		return sendError(context, "cant find category by id")
	}
	return context.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": result,
	})
}

func addCategory(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /addCategory POST")
	}

	category := &database.Category{} // Use name, icon, description
	err := context.Bind(category)
	if err != nil {
		return sendError(context, "no user information in JSON /addCategory POST")
	}

	if category.Name == "" || category.Icon == "" || category.Description == "" {
		return sendError(context, "empty params /addCategory POST")
	}

	category, err = category.Create(
		category.Name,
		category.Icon,
		category.Description,
	)

	if err != nil {
		return sendError(context, "category not created /addCategory POST")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":      "success",
		"category_id": category.ID,
	})

}

func setCategory(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /addCategory POST")
	}

	category := &database.Category{}
	err := context.Bind(category)
	if err != nil {
		return sendError(context, "no user information in JSON /addCategory POST")
	}

	err = category.UpdateNotAll()
	if err != nil {
		return sendError(context, "no user information in JSON /addCategory POST")
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

	uuidParam := context.Param("id")

	category := &database.Category{
		ID: uuidParam,
	}

	err := category.Delete()
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
