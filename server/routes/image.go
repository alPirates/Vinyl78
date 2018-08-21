package routes

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// PathToImages - path where will contain images
const (
	PathToImages = "../media/"
)

func addCaruselImage(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /addCaruselImage POST")
	}

	name, err := upload(context)
	if err != nil {
		return sendError(context, "can't upload image /addCaruselImage POST")
	}

	count, err := database.GetImagesCount(-1)
	if err != nil {
		return sendError(context, "can't count images /addCaruselImage POST")
	}

	_, err = (&database.Image{}).Create(
		name,
		-1,
		uint(count+1),
	)
	if err != nil {
		return sendError(context, "image not created /addCaruselImage POST")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}

func deleteCaruselImage(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /deleteCaruselImage DELETE")
	}

	idParam := context.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return sendError(context, "skip is not uint /deleteCaruselImage DELETE")
	}
	idUint := uint(id)

	err = (&database.Image{
		ID: idUint,
	}).Delete()

	if err != nil {
		return sendError(context, "can't delete from db /deleteCaruselImage DELETE")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}

func setCaruselImage(context echo.Context) error {
	return nil
}

func addImage(context echo.Context) error {

	// check token first
	// token := context.Get("token").(*jwt.Token)
	// claims := token.Claims.(*jwtUserClaims)

	// if claims.Role == 0 {
	// 	return sendError(context, "not admin /addImage POST")
	// }
	linked := context.FormValue("linked")
	linkedNumber, err := strconv.Atoi(linked)
	if err != nil {
		return sendError(context, "linked cant convert to number")
	}
	image := database.Image{}
	image.Linked = linkedNumber

	if image.Linked == 0 {
		return sendError(context, "empty params /addImage POST")
	}

	name, err := upload(context)
	if err != nil {
		return sendError(context, "can't upload image /addImage POST")
	}

	count, err := database.GetImagesCount(image.Linked)
	if err != nil {
		return sendError(context, "can't count images /addImage POST")
	}

	img := new(database.Image)
	img, err = (&database.Image{}).Create(
		name,
		image.Linked,
		uint(count+1),
	)
	if err != nil {
		return sendError(context, "image not created /addImage POST")
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"image":  img.ID,
	})

}

func deleteImage(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /deleteImage DELETE")
	}

	idParam := context.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return sendError(context, "skip is not uint /deleteImage DELETE")
	}
	idUint := uint(id)

	err = (&database.Image{
		ID: idUint,
	}).Delete()

	if err != nil {
		return sendError(context, "can't delete from db /deleteImage DELETE")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}

func setImage(context echo.Context) error {
	return nil
}

func upload(context echo.Context) (string, error) {

	file, err := context.FormFile("file")
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	mas := strings.Split(file.Filename, ".")
	resolution := mas[len(mas)-1]
	name := mas[0]

	name += fmt.Sprint(time.Now())
	nameBytes := md5.Sum([]byte(name))
	newName := fmt.Sprintf("%x", nameBytes)
	newName += "." + resolution

	stream, err := os.Create(PathToImages + newName)
	if err != nil {
		return "", err
	}
	defer stream.Close()

	if _, err = io.Copy(stream, src); err != nil {
		return "", err
	}

	return newName, err
}
