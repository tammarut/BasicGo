package api

import (
	echo "github.com/labstack/echo/v4"
	"github.com/tammarut/basicGo/amateur/echoYoutube/api/handlers"
)

func AdminGroup(a *echo.Group) {
	a.GET("/main", handlers.Admin) //=> Admin main page
}
