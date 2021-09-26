package app

import (
	"mrobles_app/routes"

	"github.com/labstack/echo/v4"
)

// API configuration
func RunApp() {
	e := echo.New()

	e.GET("/games", routes.GetGames)
	e.GET("/games/:id", routes.GetGame)
	e.Logger.Fatal(e.Start(":5000"))
}
