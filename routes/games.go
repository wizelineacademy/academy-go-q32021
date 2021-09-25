package routes

import (
	"mrobles_app/common"
	"mrobles_app/services"
	"strconv"

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

func GetGame(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := map[string]string{"message": "Param {id} must be numeric"}
		return c.JSON(400, response)
	}

	game, err := services.FindGame(id)
	if err != nil {
		response := map[string]string{"message": "Game not found"}
		return c.JSON(404, response)
	} else if game == (common.Game{}) {
		response := map[string]string{"message": "Game not found"}
		return c.JSON(404, response)
	}

	response := map[string]common.Game{"game": game}
	return c.JSON(200, response)
}
