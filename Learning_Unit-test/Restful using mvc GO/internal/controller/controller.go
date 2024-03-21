package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"


	"terragrit.io/sample/models"
)



//getAlbums responds with the list of all the albums as JSON.
func getAlbums(c *gin.Context){
	c.JSON(http.StatusOK, models.albums)
}

//postAlbums add an album from JSON received in the request body
func postAlbums(c *gin.Context){
	var newAlbum album

	//Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err !=nil{
		return
	}

	//Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

//getAlbumByID locates the album whose ID value matches the id parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context){
	id:= c.Param("id")

	//Loop over the list of albums
	for _, a := range albums{
		if a.ID == id{
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message":"album not found"})
}