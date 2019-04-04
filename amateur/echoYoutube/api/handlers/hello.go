package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

//Hello is default page; return Hello
func Hello(c echo.Context) error { //=> Default
	return c.String(http.StatusOK, "Hello, echoYoutube!")
}
