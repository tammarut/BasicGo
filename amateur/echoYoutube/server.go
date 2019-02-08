package main

//! Golang API(echo) =>> learn from Youtube
import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handler
func hello(c echo.Context) error { //=> Default
	return c.String(http.StatusOK, "Hello, echoYoutube!")
}

func getCats(c echo.Context) error {
	mobile := c.QueryParam("mobile")   //=> receive from param
	storage := c.QueryParam("storage") //=> receive from param
	typee := c.QueryParam("type")      //=> receive from param

	dataType := c.Param("data") //=> what datatype you want [string, json]
	if dataType == "string" {   //=> if 'string': return c.String
		return c.String(http.StatusOK, fmt.Sprintf("Mobile: %s\nstorage: %s\ntype: %s\n", mobile, storage, typee))
	}
	if dataType == "json" { //=> if 'json': return map[string] string
		return c.JSON(http.StatusOK, map[string]string{
			"mobile":  mobile,
			"storage": storage,
			"type":    typee,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{ //=> return Error Unknown type
		"Error": "We don't know you want Json or string",
	})
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)             //=> Hello
	e.GET("/cats/:data", getCats) //=> GET by parameter

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
