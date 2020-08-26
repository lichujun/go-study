package mySort

func InsertionSort(arr []int) {
	length := len(arr)
	InsertionIntervalSort(arr, 0, length-1)
}

func InsertionIntervalSort(arr []int, left int, right int) {
	for i := left + 1; i <= right; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}
