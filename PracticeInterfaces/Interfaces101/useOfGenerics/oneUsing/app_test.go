package oneusing_test

import (
	"testing"
	oneusing "useGenerics/oneUsing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	ivals := []int{23, 6, 4, 22, 60}
	expected := 60

	res, err := oneusing.Max(ivals)
	assert.Equal(t, expected, res)
	assert.NoError(t, err)

	fvals := []float64{23, 22, 123, 44}
	expectedFloat := 123.0

	result, err := oneusing.Max(fvals)

	assert.Equal(t, expectedFloat, result)
	assert.NoError(t, err)
}
