package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Hamster struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

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

func addCat(c echo.Context) error {
	cat := Cat{}
	defer c.Request().Body.Close()

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Error reading the request body for addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(body, &cat)
	if err != nil {
		log.Printf("Error unmarshalling in addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("This is your cat: %#v", cat)
	return c.String(http.StatusOK, "We got your cat!")
}

func addDog(c echo.Context) error {
	dog := Dog{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Error processing addDog requests: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("This is your dog: %#v", dog)
	return c.String(http.StatusOK, "We got your dog!")
}

func addHamster(c echo.Context) error {
	hamster := Hamster{}
	defer c.Request().Body.Close()

	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Error processing addHamster requests: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("This is your hamster: %#v", hamster)
	return c.String(http.StatusOK, "We got your hamster!")
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "You are on the secret admin main page")
}

func mainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "You are on the secret cookie main page")
}

func login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// TODO: Check username and password against DB after hasing the password
	if username == "Zuong" && password == "123" {
		cookie := &http.Cookie{}

		// This is the same
		// cookie := new(http.Cookie)

		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		// Cookie expires after 48 hours
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		return c.String(http.StatusOK, "You were logged in!")
	}

	return c.String(http.StatusUnauthorized, "Your username or password were wrong")
}

// -------------------------------- MIDDLEWARES --------------------------------
// Add to any response from the server and the server name
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "BlueBot/1.0")
		c.Response().Header().Set("NotReallyHeader", "thisHaveNoMeaning")
		return next(c)
	}
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")

		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "You don't have the right cookie")
			}
			log.Print(err)
			return err
		}

		if cookie.Value == "some_string" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "You don't have the right cookie")
	}
}

func main() {
	fmt.Println("Welcome to my server")

	e := echo.New()

	e.Use(ServerHeader)

	// -------------------------------- GROUP --------------------------------
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")

	// -------------------------------- MIDDLEWARES --------------------------------
	// This logs the server interaction
	adminGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	// Basic Auth Middleware
	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// check in the DB
		if username == "Zuong" && password == "123" {
			return true, nil
		}
		return false, nil
	}))

	cookieGroup.Use(checkCookie)

	cookieGroup.GET("/main", mainCookie)

	adminGroup.GET("/main", mainAdmin)

	e.GET("/login", login)

	// -------------------------------- CRUD --------------------------------
	e.GET("/", hello)
	e.GET("/cats/:dataType", getCats)
	e.POST("/cats", addCat)         // First method: json.Unmarshal - Fastest
	e.POST("/dogs", addDog)         // Second method: json.Decoder - Second fastest: preferred
	e.POST("/hamsters", addHamster) // Third method: c.Bind - slowest + relies on third party

	e.Start(":4000")
}
