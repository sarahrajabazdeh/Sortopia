def QuickSort(arr):
    if len(arr) <= 1:
        return arr
    pivot =arr[0]
    less = [i for i in arr[1:] if i <= pivot]
    greater =[i for i in arr[1:] if i > pivot]
    return QuickSort(less) + [pivot] + QuickSort(greater)
def main():
        arr = [64, 34, 25, 12, 22, 11, 90]
        qs = QuickSort(arr)
        print(qs)
if __name__ == "__main__":
    main()                    