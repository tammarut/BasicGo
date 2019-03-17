package middleware

import (
	"fmt"
	"net/http"
	"strings"

	echo "github.com/labstack/echo/v4"
)

func SetCookieMiddlewares(c *echo.Group) {
	c.Use(checkCookie)
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc { //=> Check user's cookie
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") { //=> If error likes this'named...'
				return c.String(http.StatusUnauthorized, "You don't have any cookies.") //=>meaning
			}
			fmt.Print("Here's =>", err)
			return err
		}
		if cookie.Value == "some-string" {
			return next(c)
		}
		return c.String(http.StatusUnauthorized, "Uncorrect the right cookie")
	}
}
