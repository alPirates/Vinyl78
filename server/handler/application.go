package handler

import (
	"fmt"
	"net/http"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type propertiesForApplications struct {
	Skip  uint `json:"skip" form:"skip" query:"skip"`
	Limit uint `json:"limit" form:"limit" query:"limit"`
}

type id struct {
	ID uint `json:"id" form:"id" query:"id"`
}

func addApplication(context echo.Context) error {

	applicationInfo := database.Application{}
	err := context.Bind(&applicationInfo)
	if err != nil {
		return sendError(context, "no user information in JSON /application PUT")
	}

	(&database.Application{}).Create(
		applicationInfo.Name,
		applicationInfo.Phone,
		applicationInfo.Email,
		applicationInfo.Message,
	)

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})

}

func setApplication(context echo.Context) error {
	return nil
}

func deleteApplication(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /application GET")
	}

	idApplication := id{}
	err := context.Bind(&idApplication)
	if err != nil {
		return sendError(context, "no properties information in JSON /application GET")
	}

	application := &database.Application{}
	application.Model.ID = idApplication.ID
	application.Delete()

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

	propertiesForApplication := propertiesForApplications{}
	err := context.Bind(&propertiesForApplication)
	if err != nil {
		return sendError(context, "no properties information in JSON /application GET")
	}

	applications := database.GetApplications(
		propertiesForApplication.Skip,
		propertiesForApplication.Limit,
	)

	return context.JSON(http.StatusOK, map[string]string{
		"result": fmt.Sprint(applications),
		"status": "success",
	})
}
