package main

import (
	"first-app/utils"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/hello", sayHallo)

	port := os.Getenv("PORT")

	if port == "" {
		port = "2121"
	}

	e.Logger.Fatal(e.Start(":" + port))
}

func sayHallo(c echo.Context) error {
	return c.JSON(200, map[string]any{
		"message":  "hallo dunia tipu2 aaaaaaa mudahkan ya Allah",
		"app name": utils.GetConfig("APP_NAME"),
		"db name":  utils.GetConfig("DB_NAME"),
	})
}
