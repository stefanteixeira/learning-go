package stack

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestEmptyStack(t *testing.T) {
  stack := Stack{}

  assert.True(t, stack.Empty())
}

func TestStackPush(t *testing.T) {
  stack := Stack{}

  stack.Push(1)
  stack.Push(3)
  stack.Push(5)

  assert.Equal(t, stack.Size(), 3)
}

func TestStackPop(t *testing.T) {
  stack := Stack{}

  stack.Push(1)
  stack.Push(2)
  stack.Push(3)
  stack.Pop()
  stack.Pop()

  assert.Equal(t, stack.Size(), 1)
}

func TestPopEmptyStack(t *testing.T) {
  stack := Stack{}

  _, err := stack.Pop()

  if assert.Error(t, err, "An error occurred") {
	   assert.Equal(t, err.Error(), "Empty stack!")
  }
}
