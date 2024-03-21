package main

import (
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
	"terragrit.io/sample/internal/controller"
)

//album represents data about a record album
type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64`json:"price"`
}

//albums slice to seed record album data
var albums = []album{
	{ID:"1", Title:"Blue Train", Artist:"John Coltrane", Price:56.99},
	{ID:"2", Title: "Slim Shady", Artist:"Eminem", Price:49.99},

}

func main(){
	router := gin.Default()
	router.GET("/albums", controller.getAlbums)
	router.POST("/albums", controller.postAlbums)
	router.GET("/albums/:id", controller.getAlbumByID)

	// defer func(){
	// 	if err := recover(); err !=nil{
	// 		//Hadle the error and log it
	// 		log.Println("Error running router:", err)
	// 	}
	// }()

	// #nosec
	router.Run("0.0.0.0:8080")
}