package models


type Album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var Albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Slim Shady", Artist: "Eminem", Price: 49.99},
}

func GetAlbums() []Album{

	return Albums
}
