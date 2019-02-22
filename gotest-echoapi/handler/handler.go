package handler

//! Golang Unit-test API(echo) =>>official cookbook
//! But Fail \n some what?
import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	User struct {
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
	}
	handler struct {
		db map[string]*User
	}
)

func (h *handler) createUser(c echo.Context) error { //=> Regular POST
	u := new(User) //=> ready to get new user
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func (h *handler) getUser(c echo.Context) error { //=> GET by parameter
	email := c.Param("name") //=>by name
	user := h.db[email]      //=> search in DB by map
	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)
}
