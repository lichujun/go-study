package mySort

func HeapSort(array []int) {
	arrayLength := len(array)
	for i := arrayLength / 2; i >= 0; i-- {
		sink(array, i, arrayLength)
	}
	for i := arrayLength - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		sink(array, 0, i)
	}
}

func sink(array []int, i int, n int) {
	var j int
	for 2*i+1 < n {
		j = 2*i + 1
		if j < n-1 && array[j] < array[j+1] {
			j++
		}
		if array[i] > array[j] {
			break
		}
		array[i], array[j] = array[j], array[i]
		i = j
	}
}
