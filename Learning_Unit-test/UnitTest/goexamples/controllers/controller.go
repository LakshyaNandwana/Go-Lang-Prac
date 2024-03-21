package controller

import (
	"net/http"


	"github.com/gin-gonic/gin"
	
	models "terragrit.io/sample/pkg"
)


func GetAlbumData(c *gin.Context){
	c.JSON(http.StatusOK, models.GetAlbums())
}


func PostAlbums(c *gin.Context){

	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	models.Albums = append(models.Albums, newAlbum)

	c.JSON(http.StatusCreated, newAlbum)

}


func GetAlbumsByID(c *gin.Context){
	id := c.Param("id")

	for _, a := range models.Albums{
		if a.ID == id{
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message":"Albums not found"})
}