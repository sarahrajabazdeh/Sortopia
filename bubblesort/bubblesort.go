package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	// Perform n-1 iterations
	for i := 0; i < n-1; i++ {
		// Compare adjacent elements and swap if necessary
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap elements
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// Unsorted array
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	// Sort the array using bubble sort
	bubbleSort(arr)

	// Print the sorted array
	fmt.Println("Sorted array:", arr)
}
