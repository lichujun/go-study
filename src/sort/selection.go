package mySort

func SelectSort(values []int) {
	length := len(values)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		min := i // 初始的最小值位置从0开始，依次向右
		for j := i + 1; j < length; j++ {
			if values[j] < values[min] {
				min = j
			}
		}
		// 把每次找出来的最小值与之前的最小值做交换
		values[i], values[min] = values[min], values[i]
	}
}
