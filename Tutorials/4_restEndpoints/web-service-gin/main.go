package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default() // Initialize a Gin router using Default.
	router.GET("/albums", getAlbums) // Use the GET function to associate the GET HTTP method and /albums path with a handler function (getAlbums).
	router.Run("localhost:8080") // Use the Run function to attach the router to an http.Server and start the server.
}

// getAlbums responds with the list of all albums as JSON.
//  creates JSON from the slice of album structs, writing the JSON into the response.
func getAlbums(c *gin.Context) { // takes a gin.Context parameter. It carries request details, validates and serializes JSON, and more.
	c.IndentedJSON(http.StatusOK, albums) //serialize the struct (albums) into JSON and add it to the response.
}