package handler

//! Golang Unit-test API(echo) =>>official cookbook
//! But Fail \n some what?
import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	//User is struct for receive request
	User struct {
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
	}
	handler struct {
		db map[string]*User
	}
)

func (h *handler) createUser(c echo.Context) error { //=> Regular POST
	u := new(User)                    //=> ready to get new user
	if err := c.Bind(u); err != nil { //=>receive from body
		return err
	}
	fmt.Printf("%+v\n", u)
	return c.JSON(http.StatusCreated, u) //=>return Status and ของ
}

func (h *handler) getUser(c echo.Context) error { //=> GET by parameter
	email := c.Param("name") //=>receive name
	user := h.db[email]      //=> search email in DB by map[string]*User
	if user == nil {         //=>handle when empty
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user) //=>return Status and ของ
}
