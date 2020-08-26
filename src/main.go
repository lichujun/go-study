package main

import (
	"fmt"
	"go-study/src/leetcode"
)

func main() {
	s := "abcabc"
	fmt.Println(leetcode.LengthOfLongestSubstring(s))

	/*var test []int
	test = append(test, 1, 2, 3)
	fmt.Println(leetcode.Permute(test))*/
	/*var coins []int
	coins = append(coins, 186, 419, 83, 408)
	fmt.Println(leetcode.CoinChange(coins, 6249))*/
	/*arr := random.GetRandomArray(20, 100)
	mySort.HeapSort(arr)
	fmt.Println(arr)
	arr = random.GetRandomArray(2000000, 10000000)
	startTime := time.Now()
	mySort.QuickSort(arr)
	fmt.Println(time.Since(startTime))*/
}
