package models


//album represent data about a record album
type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}



func returnListAlbum() *album{

	//albums slice to sav reecord album data

var albums = []album{
	{ID:"1", Title:"Blue Train", Artist:"John Coltrane", Price: 59.99},
	{ID:"2", Title:"Slim Shady", Artist: "Eminem", Price: 59.99},
}
	return albums
}



