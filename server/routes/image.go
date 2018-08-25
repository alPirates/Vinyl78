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
	DefaultImage = "default.png"
)

func getFileImage(context echo.Context) error {

	name := context.Param("name")

	_, err := os.Open(PathToImages + name)
	if err != nil {
		return context.File(PathToImages + DefaultImage)
	}

	return context.File(PathToImages + name)

}

func getImage(context echo.Context) error {

	id := context.QueryParams().Get("linked_id")

	images, err := database.GetImages(id)
	if err != nil {
		return sendError(context, "no images /image GET")
	}

	return context.JSON(200, images)
}

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

	type PutForm struct {
		Images []*database.Image `json:"images"`
	}
	var p PutForm
	err := context.Bind(&p)
	if err != nil {
		return sendError(context, "can't unmarshal /image PUT")
	}

	for _, imageR := range p.Images {
		imageR.UpdateNotAll()
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})

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
