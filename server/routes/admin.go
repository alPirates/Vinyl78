package routes

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func setToken(context echo.Context) error {
	return nil
}

func addCaruselImage(context echo.Context) {

}

func deleteCaruselImage(context echo.Context) {

}

func setCaruselImage(context echo.Context) {

}

func upload(context echo.Context) error {

	file, err := context.FormFile("file")
	if err != nil {
		return sendError(context, "")
	}
	src, err := file.Open()
	if err != nil {
		return sendError(context, "")
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return sendError(context, "")
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return sendError(context, "")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})

}
