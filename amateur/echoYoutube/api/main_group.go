package api

import (
	echo "github.com/labstack/echo/v4"
	"github.com/tammarut/basicGo/amateur/echoYoutube/api/handlers"
)

func MainGroup(e *echo.Echo) {
	// Routes
	e.GET("/login", handlers.LoginMember)       //=> GET login by parameter
	e.GET("/", handlers.Hello)                  //=> Hello
	e.GET("/mobile/:data", handlers.GetMobiles) //=> GET by parameter

	e.POST("/mobile", handlers.AddMobile) //=>POST by ReadAll body and Unmarashal
	e.POST("/laptop", handlers.AddLaptop) //=>POST by NewDecoder and Decode
	e.POST("/tablet", handlers.AddTablet) //=>POST by Bind
}
