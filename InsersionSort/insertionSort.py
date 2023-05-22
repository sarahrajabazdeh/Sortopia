def insertion_sort(arr):
    for i in range(1, len(arr)):
        key = arr[i]  # Select the next element to be inserted
        j = i - 1

        # Move elements of arr[0..i-1], that are greater than the key, to one position ahead of their current position
        while j >= 0 and arr[j] > key:
            arr[j + 1] = arr[j]  # Shift elements to the right
            j -= 1

        arr[j + 1] = key  # Insert the key into its correct position


# Example usage
arr = [12, 11, 13, 5, 6]
print("Unsorted array:", arr)

insertion_sort(arr)
print("Sorted array:", arr)
