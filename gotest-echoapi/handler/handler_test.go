package handler

//! 2 testcase but Fail expect "..." actual "...\n"
import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	//=> mockDB
	mockDB = map[string]*User{
		"Jon": &User{"Jon", "jon@labstack.com"},
	}

	//=>Expect
	userJSON = `{"name":"Jon","email":"jon@labstack.com"}`
)

func TestCreateUser(t *testing.T) {
	//. Setup
	e := echo.New()                                                               //=>instance e
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON)) //=>request POST "/" simulate payload(userJSON)
	//req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Content-Type", "application/json") //=>Header
	rec := httptest.NewRecorder()                      //=>initial response
	c := e.NewContext(req, rec)                        //=>initial context
	h := &handler{mockDB}                              //=> ready to shoot

	//. Assertions
	if assert.NoError(t, h.createUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)     //=>compare status and code
		assert.Equal(t, userJSON+"\n", rec.Body.String()) //=>compare "Want" and "mockDB"
	}
}

// func TestGetUser(t *testing.T) {
// 	// Setup
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetPath("/users/:name")
// 	c.SetParamNames("name")
// 	c.SetParamValues("Jon")
// 	h := &handler{mockDB}

// 	// Assertions
// 	if assert.NoError(t, h.getUser(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.Equal(t, userJSON, rec.Body.String())
// 	}
// }
