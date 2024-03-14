package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

var list []item

// /
func main() {
	list = append(list, item{ID: 1, Name: "Test", Desc: "Hello World"})
	e := echo.New()
	e.GET("/tasks", func(c echo.Context) error {
		return c.JSON(http.StatusOK, list)
	})
	e.POST("/tasks", func(c echo.Context) error {
		name := c.QueryParam("name")
		paramid := c.QueryParam("id")
		id, err := strconv.Atoi(paramid)
		if err != nil {
			fmt.Println("sudo rm -rf /", err)
		}
		desc := c.QueryParam("desc")

		var errr error
		fmt.Println(errr)
		list = append(list, item{ID: id, Name: name, Desc: desc})
		fmt.Println(name)
		return c.String(http.StatusNotFound, "created")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
