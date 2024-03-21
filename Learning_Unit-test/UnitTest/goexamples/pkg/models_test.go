package models

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetAlbums(t *testing.T){
	
	actualResponse := GetAlbums()

	assert.Equal(t, 2, len(actualResponse))
}