package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func Test_getAlbums(t *testing.T) {
	mockResponse := `[{"id": "1","title": "Blue Train","artist": "John Coltrane","price": 56.99},{"id": "2","title": "Slim Shady","artist": "Eminem","price": 49.99}]`

	r := SetUpRouter()
	r.GET("/albums", getAlbums)
	req, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_postAlbums(t *testing.T) {
	mockResponse := `{"id": "3","title": "Practice","artist": "Drake","price": 49.99}`
	r := SetUpRouter()

	r.POST("/albums", postAlbums)
	testAlbum := album{
		ID:     "3",
		Title:  "Practice",
		Artist: "Drake",
		Price:  49.99,
	}
	jsonValue, _ := json.Marshal(testAlbum)
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusCreated, w.Code)
}

func Test_getAlbumByIDPass(t *testing.T) {
	mockResponse := `{"id": "2","title": "Slim Shady","artist": "Eminem","price": 49.99}`
	r := SetUpRouter()

	r.GET("/albums/:id", getAlbumByID)
	req, _ := http.NewRequest("GET", "/albums/2", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}


func Test_getAlbumsIDFail( t *testing.T){

	mockResponse :=`{"Message": "album not found"}`
	r := SetUpRouter()

	r.GET("/albums/:id", getAlbumByID)
	req, _ := http.NewRequest("GET", "/albums/5", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusNotFound, w.Code)

}