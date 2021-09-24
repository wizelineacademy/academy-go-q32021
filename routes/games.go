package routes

import (
	"mrobles_app/common"
	"mrobles_app/services"

	"github.com/labstack/echo/v4"
)

func GetGames(c echo.Context) error {
	games, err := services.FindGames()
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, map[string][]common.Game{
		"games": games,
	})
}
