package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type Mobilest struct {
	Mobile  string `json: "mobile, omitempty"`
	Storage string `json: "storage, omitempty"`
	Type    string `json: "type, omitempty"`
}

func GetMobiles(c echo.Context) error {
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

func AddMobile(c echo.Context) error { //=> Solution1: Fastest ReadAll and Unmarshal
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
