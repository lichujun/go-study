package mySort

const M = 6

func QuickSort(array []int) {
	sortByQuick(array, 0, len(array)-1)
}

func sortByQuick(array []int, left int, right int) {
	if right <= left+M {
		InsertionIntervalSort(array, left, right)
		return
	}
	j := partition(array, left, right)
	sortByQuick(array, left, j-1)
	sortByQuick(array, j+1, right)
}

func partition(array []int, left int, right int) int {
	value := array[left]
	i, j := left, right+1
	for true {
		// 从左侧开始查找比value大的值
		for i++; i <= right && value > array[i]; i++ {
		}
		// 从右侧开始查找比value小的值
		for j--; j > left && value < array[j]; j-- {
		}
		// 如果从左侧开始查找比value大的值的index大于从右侧开始查找比value小的值的index，就不需要进行替换
		if i >= j {
			break
		}
		array[i], array[j] = array[j], array[i]
	}
	// 将value放入到正确的位置
	array[left], array[j] = array[j], array[left]
	return j
}
