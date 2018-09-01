package main

import (
	"github.com/alPirates/Vinyl78/server/database"
	"github.com/alPirates/Vinyl78/server/routes"
	"github.com/alPirates/Vinyl78/server/tgbot"
	"github.com/labstack/echo"
)

func main() {
	go tgbot.StartBot()
	server := echo.New()

	database.OpenConnection("vinyl78")
	defer database.CloseConnection()

	routes.SetRoutes(server)
}
