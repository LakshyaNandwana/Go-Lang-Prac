package oneusinginterface_test

import (
	"testing"
	oneusinginterface "useGenerics/oneUsingInterface"

	"github.com/stretchr/testify/assert"
)

type Group int

func TestMain(t *testing.T) {

	ivals := []int{23, 6, 4, 22, 60}
	expected := 60

	res, err := oneusinginterface.Max(ivals)
	assert.Equal(t, expected, res)
	assert.NoError(t, err)

	fvals := []float64{23, 22, 123, 44}
	expectedFloat := 123.0

	result, err := oneusinginterface.Max(fvals)

	assert.Equal(t, expectedFloat, result)
	assert.NoError(t, err)

	gvals := []Group{234, 6, 4, 22}
	var expectedRes Group = 234

	resultGroup, err := oneusinginterface.Max(gvals)

	assert.Equal(t, expectedRes, resultGroup)
	assert.NoError(t, err)
}
