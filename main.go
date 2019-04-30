package main

import (
	"cheese.self/echo-learn/util"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	registerRoute(e)
	util.RegisterToConsul(e, "Test", "127.0.0.1", 8080)
	e.Logger.Debug(e.Start(":8080"))
}