package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

type (
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

// /user Post
func AddUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Add Some Logic Here

	return c.JSON(http.StatusCreated, u)
}

// /user Get
func GetUser(c echo.Context) error {
	u := new(User)
	u.ID = 1
	u.Name = "Cheese"
	return c.JSON(http.StatusOK, u)
}
