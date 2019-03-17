package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type Laptopst struct {
	Laptop string `json: "laptop, omitempty"`
	Cpu    string `json: "cpu, omitempty"`
	Ram    string `json: "ram, omitempty"`
}

func AddLaptop(c echo.Context) error { //=> Solution2: Fastest NewDecoder and Decode
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
