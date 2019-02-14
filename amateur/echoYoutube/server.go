package main

//! Golang API(echo) =>> learn from Youtube
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
	b, err := ioutil.ReadAll(c.Request().Body) //* Here's a key
	if err != nil {
		log.Println("Failed reading body from add Mobile:", err)
		return c.String(http.StatusInternalServerError, "Error =>func addMobile Request Body.")
	}

	err = json.Unmarshal(b, &mobile) //* Here's a key
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

	err := json.NewDecoder(c.Request().Body).Decode(&laptop) //* Here's key
	if err != nil {
		log.Println("Failed processing func addLaptop NewDecoder:", err)
		c.String(http.StatusInternalServerError, "Error =>func addLaptop Newdecoder")
	}

	fmt.Printf("%#v\n", laptop)
	return c.String(http.StatusOK, "We got your laptop!!")
}

func addTablet(c echo.Context) error { //=> Solution3: Slow than both solutions above.
	tablet := Tabletst{}

	err := c.Bind(&tablet) //* Here
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
func admin(c echo.Context) error { //=> Admin main page
	return c.String(http.StatusOK, "Welcome to the secret admin main page!")
}

func main() {
	// Echo instance
	e := echo.New()
	g := e.Group("/admin") //=> Group admin

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ //=> Custom log
		Format: `[${time_rfc3339}]  status=${status}  ${method} ${host}${path}  ${latency_human} ${latency}` + "\n", //=> [2019-02-11T23:56:04+07:00]  status=200  GET localhost:1323/admin/main  0s
	}))
	e.Use(middleware.Recover())
	g.Use(middleware.BasicAuth(loginAdmin)) //=> Authorization (Username:Password)

	g.GET("/main", admin) //=> Admin main page

	// Routes
	e.GET("/", hello)                  //=> Hello
	e.GET("/mobile/:data", getMobiles) //=> GET by parameter
	e.POST("/mobile", addMobile)       //=>POST by ReadAll body and Unmarashal
	e.POST("/laptop", addLaptop)       //=>POST by NewDecoder and Decode
	e.POST("/tablet", addTablet)       //=>POST by Bind
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
