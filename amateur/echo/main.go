package main

//! Learn from youtube(API golang)
//* and unfinished yet!
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// Handler
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, TodoList{Name: "Hello, World!"})
}
func getLists(c echo.Context) error {
	listName := c.QueryParam("name")
	listType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your name is: %s\nand type is: %s\n", listName, listType))
	}
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": listName,
			"type": listType,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "need to tell us if you want json or string data",
	})

}
func addList(c echo.Context) error { //POSTแบบ1 fastest
	list := TodoList{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body for addLists: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &list)
	if err != nil {
		log.Printf("Failed unmarshaling in addLists: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	log.Printf("this is your list: %#v", list)
	return c.String(http.StatusOK, "we got your list!.")

}
func addDuedate(c echo.Context) error { //POSTแบบ2 Newdecode wise
	duedate := Duedate{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&duedate)
	if err != nil {
		log.Printf("Failed processing addDuedate request %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Printf("this's your duadate: %#v", duedate)
	return c.String(http.StatusOK, "We got your duedate.")
}
func addActivity(c echo.Context) error { //POSTแบบ3 Bind slow
	act := Activities{}

	err := c.Bind(&act)
	if err != nil {
		log.Printf("Failed addActivity request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Printf("this's your activity: %#v", act)
	return c.String(http.StatusOK, "activities added.")
}
func mainPage(c echo.Context) error {
	return c.String(http.StatusOK, "This's main page.")
}
func writeJsonFile() {
	todoLists := []TodoList{
		{Name: "Going office", Duedate: "28/12/18"},
		{Name: "Work 3 hours", Duedate: "28/12/18"},
	}
	db := ListDb{Lists: todoLists, Type: "Simple"}

	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.Encode(db)

	file, err := os.Create("todolist.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	io.Copy(file, buf)
}
func main() {
	// Echo instance
	// e := echo.New()
	// g := e.Group("/admin")

	// // Middleware
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	// }))
	// e.Use(middleware.Recover())

	// // Routes
	// g.GET("/main", mainPage)
	// e.GET("/", hello)
	// e.GET("/lists/:data", getLists)

	// e.POST("/lists", addList)
	// e.POST("/duedate", addDuedate)
	// e.POST("/activities", addActivity)

	// // Start server
	// e.Logger.Fatal(e.Start(":1323"))

	writeJsonFile()
	file, err := os.Open("todolist.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)

	db := ListDb{}
	dec.Decode(&db)
	fmt.Println(db)
}

type TodoList struct {
	Name    string `json:"name"`
	Duedate string `json:"duedate"`
}
type ListDb struct {
	Lists []TodoList
	Type  string
}
type Duedate struct {
	Day   string `json:"day"`
	Month string `json:"month"`
	Year  string `json:"year"`
}
type Activities struct {
	Name string `json: "name"`
}
