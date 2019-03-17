package middleware

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetJwtMiddlerwares(j *echo.Group) {
	j.Use(middleware.JWTWithConfig(middleware.JWTConfig{ //=> Configture
		SigningMethod: "HS512",
		SigningKey:    []byte("mySecret"),
		TokenLookup:   "cookie:JWTCookie", //=> Middleware for JWT Cookie [Save time copy]
	}))
}
