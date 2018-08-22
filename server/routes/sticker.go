package routes

import (
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

	uuid := context.Param("id")

	sticker := &database.Sticker{
		ID: uuid,
	}
	err := sticker.Delete()
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

	sticker := &database.Sticker{} // Use description, categoryUUID
	err := context.Bind(sticker)
	if err != nil {
		return sendError(context, "no user information in JSON /sticker POST")
	}

	if sticker.Description == "" || sticker.CategoryID == "" {
		return sendError(context, "empty params /sticker POST")
	}

	sticker, err = sticker.Create(
		sticker.Description,
		sticker.CategoryID,
	)

	if err != nil {
		return sendError(context, "sticker not created /sticker POST")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":     "success",
		"sticker_id": sticker.ID,
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
		return sendError(context, "limit is not uint /sticker GET")
	}
	limitUint := uint(limit)

	categoryParam := context.QueryParam("category_id")

	stickers, err := database.GetStickers(
		skipUint,
		limitUint,
		categoryParam,
	)
	if err != nil {
		return sendError(context, "sticker not returned from db /sticker GET")
	}

	for _, sticker := range stickers {
		images, err1 := database.GetImages(sticker.ID)
		if err1 != nil {
			return sendError(context, "images not returned from db /sticker GET")
		}
		sticker.Images = images
	}

	count, err := database.GetStickersCount(categoryParam)
	if err != nil {
		return sendError(context, "can't get count of stickers /sticker GET")
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"result": stickers,
		"status": "success",
		"count":  count,
	})
}
