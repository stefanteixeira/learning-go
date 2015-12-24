package main

import (
  "errors"
)

type Stack struct {
  values []interface{}
}

func (stack Stack) Size() int {
  return len(stack.values)
}

func (stack Stack) Empty() bool {
  return stack.Size() == 0
}

func (stack *Stack) Push(value interface{}) {
  stack.values = append(stack.values, value)
}

func (stack *Stack) Pop() (interface{}, error) {
  if stack.Empty() {
    return nil, errors.New("Empty stack!")
  }
  value := stack.values[stack.Size()-1]
  stack.values = stack.values[:stack.Size()-1]
  return value, nil
}
