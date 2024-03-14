package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/liandeguo/crud-golang/database"
)

var list []database.Item

// /
func main() {
	db := database.DB{}
	fmt.Println(db)
	db.GetItem(1)
	list = append(list, database.Item{ID: 1, Name: "Test", Desc: "Hello World"})
	e := echo.New()
	e.GET("/tasks", func(c echo.Context) error {
		return c.JSON(http.StatusOK, list)
	})
	e.POST("/tasks", func(c echo.Context) error {
		name := c.QueryParam("name")
		paramid := c.QueryParam("id")
		id, err := strconv.Atoi(paramid)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusBadRequest, "Skill Issue")
		}

		desc := c.QueryParam("desc")

		var errr error
		fmt.Println(errr)
		list = append(list, database.Item{ID: id, Name: name, Desc: desc})
		fmt.Println(name)
		return c.String(http.StatusNotFound, "created")
	})
	e.PUT("/tasks", func(c echo.Context) error {
		return c.String(http.StatusOK, "Lets gooooo")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
