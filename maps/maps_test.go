package maps

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestMaps(t *testing.T) {
  words := []string{"go", "java", "groovy", "clojure"}
  stats := getStats(words)
  expected := map[string]int{
    "G": 2,
    "J": 1,
    "C": 1}

  assert.Equal(t, stats, expected)
}
