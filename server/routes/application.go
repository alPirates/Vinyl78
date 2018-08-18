package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func addApplication(context echo.Context) error {

	application := database.Application{} // Use name, phone, email, message
	err := context.Bind(&application)
	if err != nil {
		return sendError(context, "no user information in JSON /application POST")
	}

	if application.Name == "" || application.Phone == "" ||
		application.Email == "" || application.Message == "" {
		return sendError(context, "empty params /application POST")
	}

	_, err = application.Create(
		application.Name,
		application.Phone,
		application.Email,
		application.Message,
	)

	if err != nil {
		return sendError(context, "application not created /application POST")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})

}

func setApplication(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /application DELETE")
	}

	return sendError(context, "not work!!!!!! not use!!!!!")

}

func deleteApplication(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /application DELETE")
	}

	application := database.Application{} // Use only ID
	err := context.Bind(&application)
	if err != nil {
		return sendError(context, "no properties information in JSON /application DELETE")
	}

	err = application.Delete()

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

	request := context.Request()
	skip, err := strconv.ParseUint(request.Header.Get("skip"), 10, 64)
	if err != nil {
		return sendError(context, "skip is not uint /application GET")
	}
	skipUint := uint(skip)

	limit, err := strconv.ParseUint(request.Header.Get("limit"), 10, 64)
	if err != nil {
		return sendError(context, "limit is not uint /application GET")
	}
	limitUint := uint(limit)

	applications, err := database.GetApplications(
		skipUint,
		limitUint,
	)
	if err != nil {
		return sendError(context, "application not returned from db /application GET")
	}

	applicationsBytes, err := json.Marshal(applications)
	if err != nil {
		return sendError(context, "not convert to json /application GET")
	}

	count, err := database.GetApplicationsCount()
	if err != nil {
		return sendError(context, "can't get count of applications /application GET")
	}
	countStr := strconv.Itoa(count)

	return context.JSON(http.StatusOK, map[string]string{
		"result": string(applicationsBytes),
		"status": "success",
		"count":  countStr,
	})
}
