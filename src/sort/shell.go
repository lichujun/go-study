package mySort

func ShellSort(array []int) {
	arrayLength := len(array)
	h := 1
	for h < arrayLength/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < arrayLength; i++ {
			for j := i; j-h >= 0 && j < arrayLength && array[j] < array[j-h]; j = j - h {
				array[j], array[j-h] = array[j-h], array[j]
			}
		}
		h = h / 3
	}
}
