package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetRoute(t *testing.T){
	router := SetUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET","/ping",nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}


func TestGetAlbumsPass(t *testing.T) {
	mockResponse := `[{"id":"1","title":"Blue Train","artist":"John Coltrane","price":56.99},{"id":"2","title":"Slim Shady","artist":"Eminem","price":49.99}]`

	r := SetUpRouter()
	r.GET("/albums", getAlbums)
	req, _ := http.NewRequest("GET", "/getalbums", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAlbumsFail(t *testing.T){
	mockResponse := `[{"Hello":"World"}]`

	r := SetUpRouter()
	r.GET("/albums", getAlbums)
	req, _ := http.NewRequest("GET", "/getalbums", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.NotEqual(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}



func TestPostAlbums(t *testing.T) {
	mockResponse := `{"id": "3","title": "Practice","artist": "Drake","price": 49.99}`
	r := SetUpRouter()

	r.POST("/createalbums", postAlbums)
	testAlbum := album{
		ID:     "3",
		Title:  "Practice",
		Artist: "Drake",
		Price:  49.99,
	}
	jsonValue, _ := json.Marshal(testAlbum)
	req, _ := http.NewRequest("POST", "/postalbums", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusCreated, w.Code)
}



func TestGetAlbumsByID(t *testing.T) {
	mockResponse := `{"id": "2","title": "Slim Shady","artist": "Eminem","price": 49.99}`
	r := SetUpRouter()

	r.GET("/getalbums/:id", getAlbumByID)
	req, _ := http.NewRequest("GET", "/getalbumsbyid/2", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}


func TestGetAlbumsByIDFail( t *testing.T){

	mockResponse :=`{"Message": "album not found"}`
	r := SetUpRouter()

	r.GET("/getalbums/:id", getAlbumByID)
	req, _ := http.NewRequest("GET", "/getalbumsbyid/5", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusNotFound, w.Code)
}