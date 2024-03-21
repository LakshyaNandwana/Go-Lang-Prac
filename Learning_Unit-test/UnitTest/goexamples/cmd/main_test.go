package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	models "terragrit.io/sample/pkg"

	"github.com/stretchr/testify/assert"
)

func TestSetUpRouter(t *testing.T) {
	router := SetUpRouter()

	assert.Equal(t, "terragrit.io/sample/controllers.GetAlbumData", router.Routes()[0].Handler)
	assert.Equal(t, "GET", router.Routes()[0].Method)
	assert.Equal(t, "/get/albums", router.Routes()[0].Path)

	assert.Equal(t, "terragrit.io/sample/controllers.GetAlbumsByID", router.Routes()[1].Handler)
	assert.Equal(t, "GET", router.Routes()[1].Method)
	assert.Equal(t, "/get/album/:id", router.Routes()[1].Path)

	assert.Equal(t, "terragrit.io/sample/controllers.PostAlbums", router.Routes()[2].Handler)
	assert.Equal(t, "POST", router.Routes()[2].Method)
	assert.Equal(t, "/add/albums", router.Routes()[2].Path)
}
func TestGetAlbumsPass(t *testing.T) {
	mockResponse := `[{"id":"1","title":"Blue Train","artist":"John Coltrane","price":56.99},{"id":"2","title":"Slim Shady","artist":"Eminem","price":49.99}]`
	router := SetUpRouter()

	req, _ := http.NewRequest("GET", "/get/albums", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetAlbumsFail(t *testing.T) {
	mockResponse := `[{"Hello":"World"}]`

	router := SetUpRouter()

	req, _ := http.NewRequest("GET", "/get/albums", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)

	assert.NotEqual(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostAlbums(t *testing.T) {
	mockResponse := `{"id": "3","title": "Practice","artist": "Drake","price": 49.99}`

	r := SetUpRouter()

	mockPostAlbumData := models.Album{
		ID:     "3",
		Title:  "Practice",
		Artist: "Drake",
		Price:  49.99,
	}

	jsonValue, _ := json.Marshal(mockPostAlbumData)
	req, _ := http.NewRequest("POST", "/add/albums", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAlbumsByIDPass(t *testing.T) {
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

func TestGetAlbumsByIDFail(t *testing.T) {

	mockResponse := `{"Message": "Albums not found"}`

	r := SetUpRouter()

	req, _ := http.NewRequest("GET", "/get/album/5", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusNotFound, w.Code)
}
