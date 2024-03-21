package main

import (
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Slim Shady", Artist: "Eminem", Price: 49.99},
}

func main() {
	router := SetUpRouter()
	// #nosec
	router.Run(":8080")
}

func SetUpRouter() *gin.Engine{
	r := gin.Default()
	r.GET("/ping",func (c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/getalbums", getAlbums)

	r.POST("postalbums", postAlbums)

	r.GET("/getalbumsbyid/:id", getAlbumByID)
	return r
}


// getAlbums responds with the list of all the albums as JSON.
func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

// postAlbums add an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	//Call BindJSON to bind the received JSON to newAlbum
	_ = c.BindJSON(&newAlbum)

	//Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//Loop over the list of albums
	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "album not found"})
}
