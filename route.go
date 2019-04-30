package main

import (
	"cheese.self/echo-learn/controller"
	"github.com/labstack/echo"
)

func registerRoute(e *echo.Echo) {
	e.POST("/user", controller.AddUser)
	e.GET("/user", controller.GetUser)
}
