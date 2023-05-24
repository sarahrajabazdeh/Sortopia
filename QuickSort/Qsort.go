package main

import "fmt"

func Quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[0] // Choose the first element as the pivot
	var left, right []int

	for _, num := range arr[1:] {
		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	Quicksort(left)  // Recursively sort the left subarray
	Quicksort(right) // Recursively sort the right subarray

	copy(arr, append(append(left, pivot), right...)) // Combine the sorted subarrays and pivot
}

func main() {
	arr := []int{5, 2, 9, 1, 7, 6, 3}
	Quicksort(arr)
	fmt.Println(arr)
}
