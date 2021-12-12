package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID 		string  `json:"id"`
	Title 	string  `json:"title"`
	Artist 	string  `json:"artist"`
	Price 	float64 `json:"price"`
}

type hello struct {
	Message string `json:"message"`
}

var helloMessage = []hello{
	{
		Message : "Hello To this Server",
	},
}

var albums = []album{
	{
		ID : "1",
		Title : "Blue Train",
		Artist : "Rahul Pawar",
		Price : 50.33,
	},
	{
		ID : "2",
		Title : "Train Blue",
		Artist : "Rahul Pawar",
		Price : 89.33,
	},
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}
	router := gin.Default()

	router.GET("/", getHelloMessage)
	router.GET("/albums", getAlbums)
	
	router.Run(":"+port)
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getHelloMessage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, helloMessage)
}