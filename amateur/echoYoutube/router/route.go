package router

import (
	echo "github.com/labstack/echo/v4"
	"github.com/tammarut/basicGo/amateur/echoYoutube/api"
	"github.com/tammarut/basicGo/amateur/echoYoutube/api/middleware"
)

//New is instance
func New() *echo.Echo {
	//.Instance
	e := echo.New()

	//.Create group
	a := e.Group("/admin")  //=> Group admin: Created
	c := e.Group("/cookie") //=> Group cookie: Created
	j := e.Group("/jwt")    //=> Group JWT: Created

	//.Set middleware
	middleware.SetMainMiddleWares(e)
	middleware.SetAdminMiddlewares(a)
	middleware.SetCookieMiddlewares(c)
	middleware.SetJwtMiddlerwares(j)

	//.Set main route
	api.MainGroup(e)

	//.Set group route
	api.AdminGroup(a)
	api.CookieGroup(c)
	api.JwtGroup(j)
	return e
}
