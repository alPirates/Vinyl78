package main

import (
	"github.com/alPirates/Vinyl78/server/database"
	"github.com/alPirates/Vinyl78/server/routes"
	"github.com/labstack/echo"
)

func main() {
	server := echo.New()
	server.Debug = true

	database.OpenConnection("vinyl78")
	defer database.CloseConnection()

	routes.SetRoutes(server)
}
