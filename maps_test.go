package main

import (
  "reflect"
  "testing"
)

func TestMaps(t *testing.T) {
  words := []string{"go", "java", "groovy", "clojure"}
  stats := getStats(words)
  expected := map[string]int{
    "G": 2,
    "J": 1,
    "C": 1}

  if !reflect.DeepEqual(stats, expected) {
    t.Error("Error getting stats!")
  }
}
