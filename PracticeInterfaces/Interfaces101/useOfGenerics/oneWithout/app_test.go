package onewithout_test

import (
	"testing"
	onewithout "useGenerics/oneWithout"

	"github.com/stretchr/testify/assert"
)

func TestMaxInts(t *testing.T) {

	//Input length greater than 0
	vals := []int{23, 8, 4, 42, 16, 15}

	expected := 42

	res, err := onewithout.MaxInts(vals)

	assert.Equal(t, res, expected)

	assert.NoError(t, err)

	//Input length zero

	res, err = onewithout.MaxInts(nil)

	assert.Equal(t, res, 0)
	assert.Error(t, err)

}

func TestMaxFloat(t *testing.T) {

	//Input length greater than 0
	vals := []float64{23, 8, 4, 42, 16, 15}

	expected := 42.0

	res, err := onewithout.MaxFloat64s(vals)

	assert.Equal(t, res, expected)

	assert.NoError(t, err)

	//Input length zero

	res, err = onewithout.MaxFloat64s(nil)

	assert.Equal(t, res, 0.0)
	assert.Error(t, err)

}
