package main

import (
	"GoUnitTest/controllers"

	"github.com/gin-gonic/gin"
)

func main(){

	router := SetUpRouter()

	// #nosec
	router.Run(":8080")
}


func SetUpRouter() *gin.Engine{

	r := gin.Default()

	r.GET("/get/albums", controller.GetAlbumData)

	r.POST("/add/albums", controller.PostAlbums)

	r.GET("/get/album/:id", controller.GetAlbumsByID)

	return r
}