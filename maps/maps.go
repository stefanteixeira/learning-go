package maps

import (
  "fmt"
  "strings"
)

func getStats(words []string) map[string]int {
  stats := make(map[string]int)

  for _, word := range words {
    first := strings.ToUpper(string(word[0]))
    count, found := stats[first]

    if found {
      stats[first] = count + 1
    } else {
      stats[first] = 1
    }
  }

  return stats
}

func print(stats map[string]int) {
  for first, count := range stats {
    fmt.Printf("%s = %d\n", first, count)
  }
}
