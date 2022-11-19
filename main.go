package main

import (
	"first-app/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/hello", sayHallo)

	e.Logger.Fatal(e.Start(":2121"))
}

func sayHallo(c echo.Context) error {
	return c.JSON(200, map[string]any{
		"message":  "hallo dunia tipu2",
		"app name": utils.GetConfig("APP_NAME"),
		"db name":  utils.GetConfig("DB_NAME"),
	})
}
