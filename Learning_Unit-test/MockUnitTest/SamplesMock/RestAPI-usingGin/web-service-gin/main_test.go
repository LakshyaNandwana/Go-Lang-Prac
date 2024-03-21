package main

import (
	"testing"

	"MockUnitTest/mocks"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)




func TestGetAll(t *testing.T){
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	var mockAlbumRepo MockalbumRepoitoryMockRecorder

	defer ctrl.Finish()

	albumsMock := mock_main.NewMockalbumRepoitory(ctrl)

	expectedAlbums := []album{
        {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
        {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
        {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
    }

	albumsMock.EXPECT().GetAll().Return(expectedAlbums)

	actualAlbums := memoryAlbumRepository(albumsMock)
	assert.Equal(t, expectedAlbums,actualAlbums)
}