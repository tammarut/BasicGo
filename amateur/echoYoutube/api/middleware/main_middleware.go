package middleware

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//
func SetMainMiddleWares(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ //=> Custom log
		Format: `[${time_rfc3339}]  status=${status}  ${method} ${host}${path}  ${latency_human} ${latency}` + "\n", //=> [2019-02-11T23:56:04+07:00]  status=200  GET localhost:1323/admin/main  0s
	}))

	e.Use(middleware.Recover())
	e.Use(ServerHeader)
}

func ServerHeader(header echo.HandlerFunc) echo.HandlerFunc { //=> Set what header you need
	return func(c echo.Context) error { //. Return Anonymous func
		c.Response().Header().Set(echo.HeaderServer, "khunPleum1.0") //=> OK
		c.Response().Header().Set("SomeHeader", "Superman")          //=> Some funny
		return header(c)
	}
}
