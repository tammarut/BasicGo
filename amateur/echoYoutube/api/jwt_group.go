package api

import (
	echo "github.com/labstack/echo/v4"
	"github.com/tammarut/basicGo/amateur/echoYoutube/api/handlers"
)

func JwtGroup(j *echo.Group) {
	j.GET("/main", handlers.Jwtmain) //=> JWT main page
}
