package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func addApplication(context echo.Context) error {

	application := &database.Application{} // Use name, phone, email, message
	err := context.Bind(application)
	if err != nil {
		return sendError(context, "no user information in JSON /application POST")
	}

	if application.Name == "" || application.Phone == "" ||
		application.Email == "" || application.Message == "" {
		return sendError(context, "empty params /application POST")
	}

	application, err = application.Create(
		application.Name,
		application.Phone,
		application.Email,
		application.Message,
	)

	if err != nil {
		return sendError(context, "application not created /application POST")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":         "success",
		"application_id": application.ID,
	})

}

func setApplication(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /application PUT")
	}

	return sendError(context, "not work!!!!!! not use!!!!!")

}

func deleteApplication(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /application DELETE")
	}

	idParam := context.Param("id")
	fmt.Println(idParam)

	application := &database.Application{
		ID: idParam,
	}
	err := application.Delete()

	if err != nil {
		return sendError(context, "application not deleted /application DELETE")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}

func getApplication(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /application GET")
	}

	skipParam := context.QueryParam("skip")
	skip, err := strconv.ParseUint(skipParam, 10, 64)
	if err != nil {
		return sendError(context, "skip is not uint /application DELETE")
	}
	skipUint := uint(skip)

	limitParam := context.QueryParam("limit")
	limit, err := strconv.ParseUint(limitParam, 10, 64)
	if err != nil {
		return sendError(context, "lmit is not uint /application DELETE")
	}
	limitUint := uint(limit)

	applications, err := database.GetApplications(
		skipUint,
		limitUint,
	)
	if err != nil {
		return sendError(context, "application not returned from db /application GET")
	}

	count, err := database.GetApplicationsCount()
	if err != nil {
		return sendError(context, "can't get count of applications /application GET")
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"result": applications,
		"status": "success",
		"count":  count,
	})
}
