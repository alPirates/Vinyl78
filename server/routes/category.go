package routes

import (
	"net/http"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func getCategory(context echo.Context) error {

	categories, err := database.GetCategories()
	if err != nil {
		return sendError(context, "can't get categories", "не удалось получить категории")
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"result":  categories,
		"message": "категории получены",
	})
}

func getCategoryByID(context echo.Context) error {
	uuidParam := context.Param("id")

	if !database.CheckUUID(uuidParam) {
		return sendError(context, "incorrect id", "не удалось получить категорию")
	}

	result, err := database.GetCategoryByID(uuidParam)
	if err != nil {
		return sendError(context, "can't find category", "не удалось получить категорию")
	}
	return context.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"result":  result,
		"message": "категория получена",
	})
}

func addCategory(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin", "вы не администратор")
	}

	category := &database.Category{} // Use name, icon, description
	err := context.Bind(category)
	if err != nil {
		return sendError(context, "no category information in JSON", "не удалось создать категорию")
	}

	if category.Name == "" || category.Icon == "" {
		return sendError(context, "empty params", "не удалось создать категорию")
	}

	category, err = category.Create(
		category.Name,
		category.Icon,
		category.Description,
	)

	if err != nil {
		return sendError(context, "category not created", "не удалось создать категорию")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":      "success",
		"category_id": category.ID,
		"message":     "категория создана",
	})

}

func setCategory(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin", "вы не администратор")
	}

	category := &database.Category{}
	err := context.Bind(category)
	if err != nil {
		return sendError(context, "no category information in JSON", "не удалось изменить категорию")
	}

	if !database.CheckUUID(category.ID) {
		return sendError(context, "incorrect id", "не удалось изменить категорию")
	}

	err = category.UpdateNotAll()
	if err != nil {
		return sendError(context, "can't update category", "не удалось изменить категорию")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "категория изменена",
	})

}

func deleteCategory(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin", "вы не администратор")
	}

	uuidParam := context.Param("id")

	if !database.CheckUUID(uuidParam) {
		return sendError(context, "incorrect id", "не удалось удалить категорию")
	}

	category := &database.Category{
		ID: uuidParam,
	}

	err := category.Delete()
	if err != nil {
		return sendError(context, "category not deleted", "не удалось удалить категорию")
	}

	err = database.DeleteByCategory(category.ID)
	if err != nil {
		return sendError(context, "stickers or images not deleted", "не удалось удалить категорию")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "категория удалена",
	})

}
