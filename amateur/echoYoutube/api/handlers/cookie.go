package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func Cookie(c echo.Context) error { //=> Cookie main page
	return c.String(http.StatusOK, "Welcome cookie page!")
}
