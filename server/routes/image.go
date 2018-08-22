package routes

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
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

func addImage(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /addImage POST")
	}

	linked := context.FormValue("linked_id")

	image := database.Image{}
	image.LinkedID = linked

	if image.LinkedID == "" {
		return sendError(context, "empty params /addImage POST")
	}

	name, err := upload(context)
	if err != nil {
		fmt.Println(err)
		return sendError(context, "can't upload image /addImage POST")
	}

	count, err := database.GetImagesCount(image.LinkedID)
	if err != nil {
		return sendError(context, "can't count images /addImage POST")
	}

	img, err := (&database.Image{}).Create(
		name,
		image.LinkedID,
		uint(count+1),
	)

	if err != nil {
		return sendError(context, "image not created /addImage POST")
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"image_id": img.ID,
	})

}

func deleteImage(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /deleteImage DELETE")
	}

	id := context.Param("id")

	err := (&database.Image{
		ID: id,
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
