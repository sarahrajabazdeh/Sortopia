# Sortopia
sorting algorithms 
1- Bubble Sort
Imagine you have a tray with a bunch of colorful bubbles of different sizes floating randomly. Your goal is to arrange them in order of their size, with the smallest bubble on the left and the largest bubble on the right. However, you can only compare two bubbles at a time.

To perform bubble sort, you start from the leftmost side of the tray and compare adjacent bubbles. If the bubble on the left is larger than the bubble on the right, you swap their positions. This causes the larger bubble to "bubble up" towards the right side of the tray.

You continue this process, comparing and swapping adjacent bubbles, moving from left to right along the tray. With each iteration, the largest bubble that has been encountered so far "bubbles up" to its correct position at the rightmost side of the tray.

This process is repeated until no more bubbles need to be swapped, indicating that the tray is fully sorted. It's like repeatedly going through the tray, comparing and moving bubbles, until they find their rightful place.

The bubble sort algorithm has a worst-case time complexity of O(n^2), where n is the number of elements in the array being sorted. This means that as the size of the array increases, the time taken to sort it grows quadratically. due to its inefficient nature, bubble sort is not recommended for large datasets and is mainly used for educational purposes or when sorting small arrays