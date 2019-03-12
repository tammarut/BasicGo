package main

//! Golang API(echo) =>> learn from Youtube
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Mobilest struct {
	Mobile  string `json: "mobile, omitempty"`
	Storage string `json: "storage, omitempty"`
	Type    string `json: "type, omitempty"`
}
type Laptopst struct {
	Laptop string `json: "laptop, omitempty"`
	Cpu    string `json: "cpu, omitempty"`
	Ram    string `json: "ram, omitempty"`
}

type Tabletst struct {
	Tablet  string `json: "tablet, omitempty"`
	Storage string `json: "storage, omitempty"`
	Screen  string `json: "screen, omitempty"`
}
type JwtClaims struct { //. struct keep token
	Name string `json:"name`
	jwt.StandardClaims
}

// Handler
func hello(c echo.Context) error { //=> Default
	return c.String(http.StatusOK, "Hello, echoYoutube!")
}

func getMobiles(c echo.Context) error {
	mobile := c.QueryParam("mobile")   //=> receive from param
	storage := c.QueryParam("storage") //=> receive from param
	typee := c.QueryParam("type")      //=> receive from param

	dataType := c.Param("data") //=> what datatype you want [string, json]
	if dataType == "string" {   //=> if 'string': return c.String
		return c.String(http.StatusOK, fmt.Sprintf("Mobile: %s\nstorage: %s\ntype: %s\n", mobile, storage, typee))
	}
	if dataType == "json" { //=> if 'json': return map[string] string
		return c.JSON(http.StatusOK, map[string]string{
			"mobile":  mobile,
			"storage": storage,
			"type":    typee,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{ //=> return Error Unknown type
		"Error": "We don't know you want Json or string",
	})
}

func addMobile(c echo.Context) error { //=> Solution1: Fastest ReadAll and Unmarshal
	mobile := Mobilest{}

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body) //. Here's a key
	if err != nil {
		log.Println("Failed reading body from add Mobile:", err)
		return c.String(http.StatusInternalServerError, "Error =>func addMobile Request Body.")
	}

	err = json.Unmarshal(b, &mobile) //. Here's a key
	if err != nil {
		log.Println("Failed unmarshaling func addMobile:", err)
		return c.String(http.StatusInternalServerError, "Error =>func addMobile unmarshaling.")
	}
	fmt.Printf("%#v\n", mobile)
	return c.String(http.StatusOK, "We got your mobiles!!")

}
func addLaptop(c echo.Context) error { //=> Solution2: Fastest NewDecoder and Decode
	laptop := Laptopst{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&laptop) //. Here's key
	if err != nil {
		log.Println("Failed processing func addLaptop NewDecoder:", err)
		c.String(http.StatusInternalServerError, "Error =>func addLaptop Newdecoder")
	}

	fmt.Printf("%#v\n", laptop)
	return c.String(http.StatusOK, "We got your laptop!!")
}

func addTablet(c echo.Context) error { //=> Solution3: Slow than both solutions above.
	tablet := Tabletst{}

	err := c.Bind(&tablet) //. Here
	if err != nil {
		log.Println("Failed processing func addTablet Bind:", err)
		c.String(http.StatusInternalServerError, "Error =>func addTablet Bind")
	}
	fmt.Printf("%#v\n", tablet)
	return c.String(http.StatusOK, "We got your tablet!!")
}

func loginAdmin(username, password string, c echo.Context) (bool, error) { //=> login admin page
	//check this in DB
	if username == "admin" && password == "0000" { //=> True Go on!
		return true, nil
	}
	return false, nil //=> Nope!
}
func loginMember(c echo.Context) error { //=> login and keep cookies
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

func ServerHeader(header echo.HandlerFunc) echo.HandlerFunc { //=> Set what header you need
	return func(c echo.Context) error { //. Return Anonymous func
		c.Response().Header().Set(echo.HeaderServer, "khunPleum1.0") //=> OK
		c.Response().Header().Set("SomeHeader", "Superman")          //=> Some funny
		return header(c)
	}
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
func admin(c echo.Context) error { //=> Admin main page
	return c.String(http.StatusOK, "Welcome to the secret admin main page!")
}
func cookie(c echo.Context) error { //=> Cookie main page
	return c.String(http.StatusOK, "Welcome cookie page!")
}
func jwtmain(c echo.Context) error { //=> JWT main page
	user := c.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	fmt.Println("User Name:", claims["name"], "User ID:", claims["jti"]) //=> print log out

	return c.String(http.StatusOK, "The JWT page!")
}

func main() {
	// Echo instance
	e := echo.New()
	a := e.Group("/admin")  //=> Group admin: Created
	c := e.Group("/cookie") //=> Group cookie: Created
	j := e.Group("/jwt")    //=> Group JWT: Created

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ //=> Custom log
		Format: `[${time_rfc3339}]  status=${status}  ${method} ${host}${path}  ${latency_human} ${latency}` + "\n", //=> [2019-02-11T23:56:04+07:00]  status=200  GET localhost:1323/admin/main  0s
	}))
	e.Use(middleware.Recover())
	e.Use(ServerHeader)                                  //=> Custom Header(ServerHeader = name)
	a.Use(middleware.BasicAuth(loginAdmin))              //=> Authorization (Username:Password)
	c.Use(checkCookie)                                   //=> Check this user ว่ามีคุกกี้นี้ในระบบไหม?
	j.Use(middleware.JWTWithConfig(middleware.JWTConfig{ //=> Configture
		SigningMethod: "HS512",
		SigningKey:    []byte("mySecret"),
		TokenLookup:   "cookie:JWTCookie", //=> Middleware for JWT Cookie [Save time copy]
	}))

	a.GET("/main", admin)   //=> Admin main page
	c.GET("/main", cookie)  //=> Cookie main page
	j.GET("/main", jwtmain) //=> JWT main page

	// Routes
	e.GET("/login", loginMember)       //=> GET login by parameter
	e.GET("/", hello)                  //=> Hello
	e.GET("/mobile/:data", getMobiles) //=> GET by parameter
	e.POST("/mobile", addMobile)       //=>POST by ReadAll body and Unmarashal
	e.POST("/laptop", addLaptop)       //=>POST by NewDecoder and Decode
	e.POST("/tablet", addTablet)       //=>POST by Bind
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
