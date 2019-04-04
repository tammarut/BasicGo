package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

//Admin is admin main page
func Admin(c echo.Context) error { //=> Admin main page
	return c.String(http.StatusOK, "Welcome to the secret admin main page!")
}
