package main

import (
	"github.com/alPirates/Vinyl78/database"
	"github.com/alPirates/Vinyl78/handler"
	"github.com/labstack/echo"
)

func main() {
	server := echo.New()

	database.OpenConnection("Vinyl78")
	defer database.CloseConnection()

	handler.SetRoutes(server)
}
