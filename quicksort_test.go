package main

import (
  "reflect"
  "testing"
)

func TestQuicksort(t *testing.T) {
  numbers := []int{9, 2, 3, 5, 7, 1}
  sorted := quicksort(numbers)
  expected := []int{1, 2, 3, 5, 7, 9}

  if !reflect.DeepEqual(sorted, expected) {
    t.Error("Quicksort error!")
  }
}
