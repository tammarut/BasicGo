package api

import (
	echo "github.com/labstack/echo/v4"
	"github.com/tammarut/basicGo/amateur/echoYoutube/api/handlers"
)

func CookieGroup(c *echo.Group) {
	c.GET("/main", handlers.Cookie) //=> Cookie main page
}
