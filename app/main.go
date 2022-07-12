package main

import (
	"dapoint-api/api"
	"dapoint-api/app/modules"
	"dapoint-api/config"
	"dapoint-api/util"
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"

	//_ "dapoint-api/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Clean Hexa Sample API
// @version 1.0
// @description Berikut API yang digunakan untuk mini project
func main() {
	config := config.GetConfig()

	dbCon := util.NewConnectionDatabase(config)
	defer dbCon.CloseConnection()

	controllers := modules.RegisterModules(dbCon, config)

	e := echo.New()
	e.Use(middleware.CORS())
	handleSwag := echoSwagger.WrapHandler
	e.GET("/swagger/*", handleSwag)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "dapoint API")
	})
	api.RegistrationPath(e, controllers)

	go func() {
		var appAddress string
		fmt.Println(config.App.Env)
		if config.App.Env == "dev" {
			appAddress = "127.0.0.1"
		} else {
			appAddress = "0.0.0.0"
		}
		address := fmt.Sprintf("%s:%d", appAddress, config.App.Port)
		if err := e.Start(address); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	<-quit
}
