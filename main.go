package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
type item struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}
var list []item




func main() {
	list = append(list,item{ID: 1,Name: "Test",Desc: "Hello World"})

	router := gin.Default()
	router.GET("/list", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, list)
	})

 	router.GET("/item/:id", func(c *gin.Context) {
        c.
    })

	router.Run("localhost:8080")
}
