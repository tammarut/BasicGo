package handlers

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	echo "github.com/labstack/echo/v4"
)

type JwtClaims struct { //. struct keep token
	Name string `json:"name`
	jwt.StandardClaims
}

func LoginMember(c echo.Context) error { //=> login and keep cookies
	username := c.QueryParam("usr") //=> get parameter from
	password := c.QueryParam("pwd")

	//!check this member in Database
	if username == "zero" && password == "1234" {
		cookie := &http.Cookie{} //=> Call struct cookie
		//cookie := new(http.Cookie) // same above
		cookie.Name = "sessionID" //=> Pattern keep cookie
		cookie.Value = "some-string"
		cookie.Expires = time.Now().Add(15 * time.Minute)
		c.SetCookie(cookie) //=> Save Cookie

		//!Create jwt token
		token, err := createJwtToken() //=> Instance call function createJwtToken
		if err != nil {
			fmt.Println("Error: Creating JWT token ", err)
			return c.String(http.StatusInternalServerError, "Error! JWT token sponse")
		}
		//! JWT Cookie
		jwtCookie := &http.Cookie{}  //=> One Struct for JWT cookie (same above cookie)
		jwtCookie.Name = "JWTCookie" //=> Same name as middlware: Tokenlookup
		jwtCookie.Value = token      //=> token here!
		cookie.Expires = time.Now().Add(15 * time.Minute)
		c.SetCookie(jwtCookie) //=> Save JWT Cookie

		return c.JSON(http.StatusOK, map[string]string{ //=> return Json_map style ;Awesome!
			"message": "Logged in successful.",
			"token":   token, //=> return token to body
		})
	}
	return c.String(http.StatusUnauthorized, "Wrong username or password")
}

func createJwtToken() (string, error) { //=> Not much, Create token
	claims := JwtClaims{ //=> JwtClaim struct; above
		"zero", //=> Name
		jwt.StandardClaims{ //=> Detail
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims) //? What's these code under?
	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}
	return token, nil //=> return
}
