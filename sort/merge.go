package mySort

// 自顶向下
func MergeSortUB(array []int) {
	arrayLength := len(array)
	aux := make([]int, arrayLength)
	sortByMerge(array, aux, 0, arrayLength-1)
}

// 自底向上
func MergeSortBU(array []int) {
	arrayLength := len(array)
	aux := make([]int, arrayLength)
	for sz := 1; sz < arrayLength; sz = sz + sz {
		for left := 0; left < arrayLength-1; left += sz + sz {
			merge(array, aux, left, left+sz-1, Min(left+sz+sz-1, arrayLength-1))
		}
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func sortByMerge(array []int, aux []int, left int, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	// 将array[left:mid]数组进行排序
	sortByMerge(array, aux, left, mid)
	// 将array[mid+1:right]数组进行排序
	sortByMerge(array, aux, mid+1, right)
	merge(array, aux, left, mid, right)
}

// 将array[left:mid]和array[mid+1:right]这两个有序数组合并成一个array[left:right]有序数组
func merge(array []int, aux []int, left int, mid int, right int) {
	if mid == right {
		return
	}
	for i := left; i <= right; i++ {
		aux[i] = array[i]
	}
	k, j := left, mid+1
	// 将两个有序数组合并为一个有序数组
	for i := left; i <= right; i++ {
		if k > mid {
			array[i] = aux[j]
			j++
		} else if j > right {
			array[i] = aux[k]
			k++
		} else if aux[k] < aux[j] {
			array[i] = aux[k]
			k++
		} else {
			array[i] = aux[j]
			j++
		}
	}
}
