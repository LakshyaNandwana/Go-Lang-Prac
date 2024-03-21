package main

import(
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


type albumRepoitory interface{
    GetAll() []album
    GetByID(id string)(album,bool)
    Add(album)
}


type  memoryAlbumRepository struct{
    albums []album
}

func ( r * memoryAlbumRepository) GetAll() []album{
    return r.albums
}
func (r *memoryAlbumRepository)GetByID(id string) (album, bool){
    for _, a := range r.albums{
        if a.ID == id{
            return a, true
        }
    }
    return album{}, false
}

func (r *memoryAlbumRepository) Add(a album){
    r.albums = append(r.albums, a)
}

func main(){
    router := gin.Default()
    repo := &memoryAlbumRepository{albums: []album{
        {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
        {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
        {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
    }}

    router.GET("/albums", func(c *gin.Context){
        c.IndentedJSON(http.StatusOK, repo.GetAll())
    })

    router.POST("/albums", func(c *gin.Context){
        var newAlbum album
        if err := c.BindJSON(&newAlbum); err !=nil{
            return
        }
        repo.Add(newAlbum)
        c.IndentedJSON(http.StatusCreated, newAlbum)
    })

    router.GET("/albums/:id", func(c *gin.Context){
        id := c.Param("id")
        a, found := repo.GetByID(id)
        if found{
            c.IndentedJSON(http.StatusOK, a)
        }else{
            c.IndentedJSON(http.StatusNotFound, gin.H{"message":"album not found"})
        }
    })

    router.Run("localhost:8080")
}






/*
// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main(){
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.POST("/albums", postAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.Run("localhost:8080")
}

//getAlbums responds with the list of all the ablums as JSON.
func getAlbums(c *gin.Context){
    c.IndentedJSON(http.StatusOK, albums)
}

//postAlbums add an album from JSON received in the request body.
func postAlbums(c *gin.Context){
    var newAlbum album

    //Call BindJSON to bind the received JSON to newAlbum
    if err := c.BindJSON(&newAlbum); err !=nil{
        return
    }

    //Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

//getAlbumByID locates the album whose ID value matches the id parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context){
    id := c.Param("id")

    //Loop over the list of albums
    for _, a := range albums{
        if a.ID ==id{
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
*/