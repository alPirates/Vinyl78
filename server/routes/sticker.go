package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func setSticker(context echo.Context) error {
	return nil
}

func deleteSticker(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /sticker DELETE")
	}

	idParam := context.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return sendError(context, "skip is not uint /sticker DELETE")
	}
	idUint := uint(id)

	sticker := &database.Sticker{
		ID: idUint,
	}
	err = sticker.Delete()
	if err != nil {
		return sendError(context, "sticker not deleted /sticker DELETE")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}

func addSticker(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin /sticker POST")
	}

	sticker := database.Sticker{} // Use description, category
	err := context.Bind(&sticker)
	if err != nil {
		return sendError(context, "no user information in JSON /sticker POST")
	}

	if sticker.Description == "" || sticker.Category == 0 {
		return sendError(context, "empty params /sticker POST")
	}

	fmt.Println(sticker)

	_, err = sticker.Create(
		sticker.Description,
		sticker.Category,
	)

	if err != nil {
		return sendError(context, "sticker not created /sticker POST")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}

func getSticker(context echo.Context) error {

	skipParam := context.QueryParam("skip")
	skip, err := strconv.ParseUint(skipParam, 10, 64)
	if err != nil {
		return sendError(context, "skip is not uint /sticker GET")
	}
	skipUint := uint(skip)

	limitParam := context.QueryParam("limit")
	limit, err := strconv.ParseUint(limitParam, 10, 64)
	if err != nil {
		return sendError(context, "lmit is not uint /sticker GET")
	}
	limitUint := uint(limit)

	categoryParam := context.QueryParam("category")
	category, err := strconv.ParseUint(categoryParam, 10, 64)
	if err != nil {
		return sendError(context, "category is not uint /sticker GET")
	}
	categoryUint := uint(category)

	stickers, err := database.GetStickers(
		skipUint,
		limitUint,
		categoryUint,
	)
	if err != nil {
		return sendError(context, "sticker not returned from db /sticker GET")
	}

	for _, sticker := range stickers {
		images, errLocal := database.GetImages(int(sticker.ID))
		if errLocal != nil {
			return sendError(context, "can't get images of stickers /sticker GET")
		}
		sticker.Images = images
	}

	count, err := database.GetStickersCount(categoryUint)
	if err != nil {
		return sendError(context, "can't get count of stickers /sticker GET")
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"result": stickers,
		"status": "success",
		"count":  count,
	})
}
