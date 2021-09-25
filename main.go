package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"

	"github.com/macrojoe/academy-go-q32021/config"
	// "github.com/macrojoe/academy-go-q32021/infrastructure/datastore"
	"github.com/macrojoe/academy-go-q32021/infrastructure/router"
	"github.com/macrojoe/academy-go-q32021/registry"
)

func main() {
	config.ReadConfig()

	// db := datastore.NewDB()
	// db.LogMode(true)
	// defer db.Close()

	r := registry.NewRegistry()

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
