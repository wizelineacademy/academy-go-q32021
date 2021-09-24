package app

import (
	"mrobles_app/routes"

	"github.com/labstack/echo/v4"
)

// configuraciones y carga de dependencias de la api
func RunApp() {
	e := echo.New()

	// e.Use(e.Logger())
	// e.Use(e.Recover())

	e.GET("/games", routes.GetGames)
	e.Logger.Fatal(e.Start(":5000"))
}
