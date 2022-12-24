package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
type Book struct{
	Id int
	Title string
	Price float64

}
func main() {
	book := Book{1,"Looking For Alaska",245.00}
	r := gin.Default()
	response,err :=json.Marshal(book)
	if err != nil {
	log.Fatal(err)
	}

	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"code":http.StatusOK,
			"message":string(response),
		})
	})
	r.Run() 
}