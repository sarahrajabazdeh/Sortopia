package main

import "fmt"

// Function to perform insertion sort
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i] // Select the next element to be inserted
		j := i - 1

		// Move elements of arr[0..i-1], that are greater than the key, to one position ahead of their current position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j] // Shift elements to the right
			j--
		}
		arr[j+1] = key // Insert the key into its correct position
	}
}

func main() {
	// Example usage
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println("Unsorted array:", arr)

	insertionSort(arr)
	fmt.Println("Sorted array:", arr)
}
