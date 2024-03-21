package main

import (
	"GoUnitTest/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestMainFunc (t *testing.T){
	testServer := httptest.NewServer(SetUpRouter())
	defer testServer.Close()


	_, err := http.Get(testServer.URL + "/get/albums")
	assert.NoError(t, err)
}

func TestGetAlbumsPass( t *testing.T){
	mockResponse := `[{"id":"1","title":"Blue Train","artist":"John Coltrane","price":56.99},{"id":"2","title":"Slim Shady","artist":"Eminem","price":49.99}]`
	router := SetUpRouter()

	
	req, _ := http.NewRequest("GET", "/get/albums", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
	
}

func TestGetAlbumsFail(t *testing.T){
	mockResponse := `[{"Hello":"World"}]`

	router := SetUpRouter()


	req, _ := http.NewRequest("GET", "/get/albums", nil)
	w:= httptest.NewRecorder()

	router.ServeHTTP(w, req)
	responseData, _ :=ioutil.ReadAll(w.Body)

	assert.NotEqual(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}


func TestPostAlbums(t *testing.T){
	mockResponse := `{"id": "3","title": "Practice","artist": "Drake","price": 49.99}`

	r := SetUpRouter()


	mockPostAlbumData := models.Album{
		ID:"3",
		Title: "Practice",
		Artist: "Drake",
		Price: 49.99,
	}

	jsonValue, _ := json.Marshal(mockPostAlbumData)
	req, _ := http.NewRequest("POST", "/add/albums", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusCreated, w.Code)
}



func TestGetAlbumsByIDPass(t *testing.T){
	mockResponse := `{"id": "2","title": "Slim Shady","artist": "Eminem","price": 49.99}`

	r := SetUpRouter()

	// r.GET("/get/albums/:id", controller.GetAlbumsByID)

	req, _ := http.NewRequest("GET", "/get/album/2", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}



func TestGetAlbumsByIDFail(t *testing.T){

	mockResponse := `{"Message": "Albums not found"}`

	r := SetUpRouter()


	req, _ := http.NewRequest("GET", "/get/album/5", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusNotFound, w.Code)
}