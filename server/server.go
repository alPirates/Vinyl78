package main

import (
	"github.com/alPirates/Vinyl78/server/database"
	"github.com/alPirates/Vinyl78/server/handler"
	"github.com/labstack/echo"
)

func main() {
	server := echo.New()

	database.OpenConnection("vinyl78")
	defer database.CloseConnection()

	handler.SetRoutes(server)
}
