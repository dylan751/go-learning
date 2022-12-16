package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from the website")
}

func getCats(c echo.Context) error {
	dataType := c.Param("dataType")
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Cat name: %s\nCat type: %s", catName, catType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "You need to let us know if you want json or string data",
	})
}

func main() {
	fmt.Println("Welcome to my server")

	e := echo.New()
	e.GET("/", hello)
	e.GET("/cats/:dataType", getCats)

	e.Start(":4000")
}
