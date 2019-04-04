package handlers

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	echo "github.com/labstack/echo/v4"
)

//Jwtmain is jwt main page
func Jwtmain(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	fmt.Println("User Name:", claims["name"], "User ID:", claims["jti"]) //=> print log out

	return c.String(http.StatusOK, "The JWT page!")
}
