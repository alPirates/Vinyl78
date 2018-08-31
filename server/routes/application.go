package routes

import (
	"net/http"
	"strconv"

	"github.com/alPirates/Vinyl78/server/database"
	"github.com/alPirates/Vinyl78/server/tgbot"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func addApplication(context echo.Context) error {

	application := &database.Application{} // Use name, phone, email, message
	err := context.Bind(application)
	if err != nil {
		return sendError(context, "no application information in JSON", "не удалось создать заявку")
	}

	if application.Name == "" || application.Phone == "" ||
		application.Email == "" || application.Message == "" {
		return sendError(context, "empty params", "не удалось создать заявку")
	}

	application, err = application.Create(
		application.Name,
		application.Phone,
		application.Email,
		application.Message,
	)

	if err != nil {
		return sendError(context, "application not created", "не удалось создать заявку")
	}

	go tgbot.SendApplication(application)

	return context.JSON(http.StatusOK, map[string]string{
		"status":         "success",
		"application_id": application.ID,
		"message":        "заяка создана",
	})

}

func setApplication(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin", "вы не администратор")
	}

	application := &database.Application{}
	err := context.Bind(application)
	if err != nil {
		return sendError(context, "no application information in JSON", "не удалось изменить заявку")
	}

	if !database.CheckUUID(application.ID) {
		return sendError(context, "incorrect id", "не удалось изменить заявку")
	}

	err = application.UpdateNotAll()
	if err != nil {
		return sendError(context, "can't update application", "не удалось изменить заявку")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "заяка изменена",
	})

}

func deleteApplication(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin", "вы не администратор")
	}

	idParam := context.Param("id")

	if !database.CheckUUID(idParam) {
		return sendError(context, "incorrect id", "не удалось удалить заявку")
	}

	application := &database.Application{
		ID: idParam,
	}
	err := application.Delete()

	if err != nil {
		return sendError(context, "appkication is not deleted", "не удалось удалить заявку")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "заяка удалена",
	})
}

func getApplication(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin", "вы не администратор")
	}

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

	applications, err := database.GetApplications(
		skipUint,
		limitUint,
	)
	if err != nil {
		return sendError(context, "can't get applications", "")
	}

	count, err := database.GetApplicationsCount()
	if err != nil {
		return sendError(context, "can't count applications", "")
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"result": applications,
		"status": "success",
		"count":  count,
	})
}
