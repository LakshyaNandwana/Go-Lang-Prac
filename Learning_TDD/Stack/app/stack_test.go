package stack

import (
	"testing"

	// stack "Stack/app"

	"github.com/stretchr/testify/assert"
)

func setUp(t *testing.T) {
	CreateStack()
}
func TestNothing(t *testing.T) {
}

func Test_CreateStack_IsEmpty(t *testing.T) {

	setUp(t)

	assert.True(t, isEmpty, "Stack is empty")

}

func Test_PushStack_IsNotEmpty(t *testing.T) {

	setUp(t)

	Push(99)

	assert.False(t, isEmpty, "Stack is not empty")
}

func Test_Pop_Panics(t *testing.T) {
	setUp(t)

	// Pop()
	assert.Panics(t, func() { Pop() })
}

func Test_PushThenPop_Empty(t *testing.T) {
	setUp(t)

	Push(99)
	Pop()

	assert.True(t, isEmpty, "Stack is empty")
}

func Test_PushTwicePopOnce_NotEmpty(t *testing.T) {
	setUp(t)

	Push(99)
	Push(88)
	Pop()

	assert.False(t, isEmpty, "Stack is not empty")
}

func Test_PushXShouldPopX(t *testing.T) {
	setUp(t)

	res := Push(99)
	expected := Pop()
	assert.Equal(t, res, expected)

}

func Test_PushTwicePopTwice(t *testing.T) {
	setUp(t)

	pushValue := Push(69)
	popValue := Pop()
	assert.Equal(t, pushValue, popValue)

	pushValue = Push(123)
	popValue = Pop()
	assert.Equal(t, pushValue, popValue)

}
