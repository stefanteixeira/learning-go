package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]
	numbers := make([]int, len(input))
	
	for i, n := range input {
		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s is not valid number!\n", n)
			os.Exit(1)
		}
		
		numbers[i] = number
	}

	fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	
	n := make([]int, len(numbers))
	copy(n, numbers)

	pivotIndex := len(n)/2
	pivot := n[pivotIndex]
	n = append(n[:pivotIndex], n[pivotIndex+1:]...)

	left, right := partition(n, pivot)

	return append(append(quicksort(left), pivot), quicksort(right)...)
}

func partition(numbers []int, pivot int) (left []int, right []int) {
	for _, n := range numbers {
		if n <= pivot {
			left = append(left, n)
		} else {
			right = append(right, n)
		}
	}

	return left, right
}
