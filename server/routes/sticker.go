package routes

import (
	"net/http"
	"strconv"

	"github.com/alPirates/Vinyl78/server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func updateStickersPosition(context echo.Context) error {
	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin", "вы не администратор")
	}
	type Req struct {
		Stickers []database.Sticker `json:"stickers" form:"stickers" query:"stickers"`
	}
	req := Req{}
	err := context.Bind(&req)

	if err != nil {
		return sendError(context, "Cant unmarshall stickers array", "Не удается обновить")
	}

	for _, sticker := range req.Stickers {
		err := sticker.UpdatePosition()
		if err != nil {
			sendError(context, "Cant update", err.Error())
		}
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "Порядок изменен",
	})
}

func setSticker(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin", "вы не администратор")
	}

	sticker := &database.Sticker{}
	err := context.Bind(sticker)

	if err != nil {
		return sendError(context, "no sticker information in JSON", "не удалось изменить стикер")
	}

	if !database.CheckUUID(sticker.ID) {
		return sendError(context, "incorrect id", "не удалось изменить стикер")
	}

	err = sticker.UpdateNotAll()
	if err != nil {
		return sendError(context, "can't update sticker", "не удалось изменить стикер")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "стикер изменен",
	})
}

func deleteSticker(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin", "вы не администратор")
	}

	uuid := context.Param("id")

	if !database.CheckUUID(uuid) {
		return sendError(context, "incorrect id", "не удалось удалить стикер")
	}

	sticker := &database.Sticker{
		ID: uuid,
	}
	err := sticker.Delete()
	if err != nil {
		return sendError(context, "sticker not deleted", "не удалось удалить стикер")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "стикер удален",
	})
}

func addSticker(context echo.Context) error {

	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(*jwtUserClaims)

	if claims.Role == 0 {
		return sendError(context, "not admin", "вы не администратор")
	}

	sticker := &database.Sticker{} // Use description, categoryUUID, text
	err := context.Bind(sticker)
	if err != nil {
		return sendError(context, "no sticker information in JSON", "не удалось создать стикер")
	}

	if sticker.Description == "" || sticker.CategoryID == "" {
		return sendError(context, "empty params", "не удалось создать стикер")
	}

	sticker, err = sticker.Create(
		sticker.Description,
		sticker.Text,
		sticker.CategoryID,
	)

	if err != nil {
		return sendError(context, "sticker not created", "не удалось создать стикер")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status":     "success",
		"sticker_id": sticker.ID,
		"message":    "стикер создан",
	})
}

func getSticker(context echo.Context) error {

	skipParam := context.QueryParam("skip")
	skip, err := strconv.ParseUint(skipParam, 10, 64)
	if err != nil {
		return sendError(context, "skip is not uint", "")
	}
	skipUint := uint(skip)

	limitParam := context.QueryParam("limit")
	limit, err := strconv.ParseUint(limitParam, 10, 64)
	if err != nil {
		return sendError(context, "limit is not uint", "")
	}
	limitUint := uint(limit)

	categoryParam := context.QueryParam("category_id")

	stickers, err := database.GetStickers(
		skipUint,
		limitUint,
		categoryParam,
	)
	if err != nil {
		return sendError(context, "can't get stickers", "")
	}

	for _, sticker := range stickers {
		images, err1 := database.GetImages(sticker.ID)
		if err1 != nil {
			return sendError(context, "can't get images", "")
		}
		sticker.Images = images
	}

	count, err := database.GetStickersCount(categoryParam)
	if err != nil {
		return sendError(context, "can't count stickers", "")
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"result": stickers,
		"status": "success",
		"count":  count,
	})
}

func getStickerByID(context echo.Context) error {
	uuid := context.Param("id")
	if uuid == "" {
		return sendError(context, "cant find sticker id", "")
	}

	if !database.CheckUUID(uuid) {
		return sendError(context, "incorrect id", "не удалось удалить стикер")
	}
	sticker, err := database.GetSticker(uuid)
	if err != nil {
		return sendError(context, "can't get sticker from db", "")
	}
	// geting sticker images
	images, err1 := database.GetImages(sticker.ID)
	if err1 != nil {
		return sendError(context, "can't get images", "")
	}
	sticker.Images = images

	return context.JSON(http.StatusOK, map[string]interface{}{
		"result": sticker,
		"status": "success",
	})
}
