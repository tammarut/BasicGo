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

func addMobile(c echo.Context) error {
	mobile := Mobilest{}

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Failed reading body from add Mobile:", err)
		return c.String(http.StatusInternalServerError, "Error =>func addMobile Request Body.")
	}

	err = json.Unmarshal(b, &mobile)
	if err != nil {
		log.Println("Failed unmarshaling func addMobile:", err)
		return c.String(http.StatusInternalServerError, "Error =>func addMobile unmarshaling.")
	}
	fmt.Printf("%#v\n", mobile)
	return c.String(http.StatusOK, "We got your mobiles!!")

}
func addLaptop(c echo.Context) error {
	laptop := Laptopst{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&laptop)
	if err != nil {
		log.Println("Failed processing func addLaptop NewDecoder:", err)
		c.String(http.StatusInternalServerError, "Error =>func addLaptop Newdecoder")
	}

	fmt.Printf("%#v\n", laptop)
	return c.String(http.StatusOK, "We got your laptop!!")
}

func addTablet(c echo.Context) error {
	tablet := Tabletst{}

	err := c.Bind(&tablet)
	if err != nil {
		log.Println("Failed processing func addTablet Bind:", err)
		c.String(http.StatusInternalServerError, "Error =>func addTablet Bind")
	}
	fmt.Printf("%#v\n", tablet)
	return c.String(http.StatusOK, "We got your tablet!!")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)                  //=> Hello
	e.GET("/mobile/:data", getMobiles) //=> GET by parameter
	e.POST("/mobile", addMobile)
	e.POST("/laptop", addLaptop)
	e.POST("/tablet", addTablet)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
