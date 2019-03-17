package middleware

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetAdminMiddlewares(a *echo.Group) {
	a.Use(middleware.BasicAuth(loginAdmin))
}

func loginAdmin(username, password string, c echo.Context) (bool, error) { //=> login admin page
	//check this in DB
	if username == "admin" && password == "0000" { //=> True Go on!
		return true, nil
	}
	return false, nil //=> Nope!
}
