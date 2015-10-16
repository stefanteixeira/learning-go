package main

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
