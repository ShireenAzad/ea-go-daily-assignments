package main

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)
type Book struct{
	Id int
	Title string
	Price float64

}
//Method 2 directly passing boook without converting to string
func main() {
	book := Book{1,"Looking For Alaska",245.00}
	r := gin.Default()
	_,err :=json.Marshal(book)
	if err!= nil {
    log.Fatal(err)
	}
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200,book)
	})
	r.Run() 
}