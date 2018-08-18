package handler

import (
	"encoding/json"
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

	sticker := database.Sticker{} // Use ID
	err := context.Bind(&sticker)
	if err != nil {
		return sendError(context, "no user information in JSON /sticker DELETE")
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

	sticker := database.Sticker{} // Use description, image, category
	err := context.Bind(&sticker)
	if err != nil {
		return sendError(context, "no user information in JSON /sticker POST")
	}

	if sticker.Description == "" || sticker.Image == "" || sticker.Category == 0 {
		return sendError(context, "empty params /sticker POST")
	}

	fmt.Println(sticker)

	_, err = sticker.Create(
		sticker.Description,
		sticker.Image,
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

	request := context.Request()
	skip, err := strconv.ParseUint(request.Header.Get("skip"), 10, 64)
	if err != nil {
		return sendError(context, "skip is not uint /sticker GET")
	}
	skipUint := uint(skip)

	limit, err := strconv.ParseUint(request.Header.Get("limit"), 10, 64)
	if err != nil {
		return sendError(context, "limit is not uint /sticker GET")
	}
	limitUint := uint(limit)

	category, err := strconv.ParseUint(request.Header.Get("category"), 10, 64)
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

	stickersBytes, err := json.Marshal(stickers)
	if err != nil {
		return sendError(context, "not convert to json /sticker GET")
	}

	count, err := database.GetStickersCount(categoryUint)
	if err != nil {
		return sendError(context, "can't get count of stickers /sticker GET")
	}
	countStr := strconv.Itoa(count)

	return context.JSON(http.StatusOK, map[string]string{
		"result": string(stickersBytes),
		"status": "success",
		"count":  countStr,
	})
}
