package routes

import (
	"net/http"
	"strconv"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func addCarousel(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin", "вы не администратор")
	}

	carousel := &database.Carousel{} // Use name, phone, email, message
	err := context.Bind(carousel)
	if err != nil {
		return sendError(context, "no carousel information in JSON", "не удалось создать карусель")
	}

	if carousel.Name == "" {
		return sendError(context, "empty params", "не удалось создать карусель")
	}

	carousel, err = carousel.Create(
		carousel.Name,
	)

	if err != nil {
		return sendError(context, "carousel not created", "не удалось создать карусель")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":         "success",
		"application_id": carousel.ID,
		"message":        "карусель создана",
	})

}

func setСarousel(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin", "вы не администратор")
	}

	carousel := &database.Carousel{}
	err := context.Bind(carousel)
	if err != nil {
		return sendError(context, "no carousel information in JSON", "не удалось изменить карусель")
	}

	if !database.CheckUUID(carousel.ID) {
		return sendError(context, "incorrect id", "не удалось изменить карусель")
	}

	err = carousel.UpdateNotAll()
	if err != nil {
		return sendError(context, "can't update carousel", "не удалось изменить карусель")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "каруслель изменена",
	})

}

func deleteCarousel(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin", "вы не администратор")
	}

	idParam := context.Param("id")

	if !database.CheckUUID(idParam) {
		return sendError(context, "incorrect id", "не удалось удалить карусель")
	}

	carousel := &database.Carousel{
		ID: idParam,
	}
	err := carousel.Delete()

	if err != nil {
		return sendError(context, "carousel is not deleted", "не удалось удалить карусель")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "карусель удалена",
	})
}

func getCarousel(context echo.Context) error {

	skipParam := context.QueryParam("skip")
	skip, err := strconv.ParseUint(skipParam, 10, 64)
	if err != nil {
		return sendError(context, "skip is not uint", "")
	}
	skipUint := uint(skip)

	limitParam := context.QueryParam("limit")
	limit, err := strconv.ParseUint(limitParam, 10, 64)
	if err != nil {
		return sendError(context, "limit is not uint", "")
	}
	limitUint := uint(limit)

	carousels, err := database.GetCarousel(
		skipUint,
		limitUint,
	)
	if err != nil {
		return sendError(context, "can't get carousel", "")
	}

	count, err := database.GetCarouselCount()
	if err != nil {
		return sendError(context, "can't count carousel", "")
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"result": carousels,
		"status": "success",
		"count":  count,
	})
}
