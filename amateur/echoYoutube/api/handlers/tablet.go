package handlers

import (
	"fmt"
	"log"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type Tabletst struct {
	Tablet  string `json: "tablet, omitempty"`
	Storage string `json: "storage, omitempty"`
	Screen  string `json: "screen, omitempty"`
}

func AddTablet(c echo.Context) error { //=> Solution3: Slow than both solutions above.
	tablet := Tabletst{}

	err := c.Bind(&tablet) //. Here
	if err != nil {
		log.Println("Failed processing func addTablet Bind:", err)
		c.String(http.StatusInternalServerError, "Error =>func addTablet Bind")
	}
	fmt.Printf("%#v\n", tablet)
	return c.String(http.StatusOK, "We got your tablet!!")
}
