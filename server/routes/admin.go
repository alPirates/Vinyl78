package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func setToken(context echo.Context) error {
	return context.String(http.StatusOK, "This can set TelegToken")
}
