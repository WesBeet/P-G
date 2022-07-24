package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type oralb struct {
	ID      string  `json:"id"`
	Product string  `json:"Product"`
	Series  string  `json:"Series"`
	Price   float64 `json:"price"`
}

var Oralb = []oralb{
	{ID: "1", Product: "Oral-B IO", Series: "6", Price: 149.99},
	{ID: "2", Product: "Oral-B iO", Series: "7G", Price: 199.99},
	{ID: "3", Product: "Oral-B iO", Series: "8", Price: 249.99},
	{ID: "4", Product: "Oral-B iO", Series: "9", Price: 299.99},
}

func getOralb(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Oralb)
}

func main() {
	router := gin.Default()
	router.GET("/oralb", getOralb)

	router.Run("localhost:8080")
}
