package random

import (
	"math/rand"
	"time"
)

func GetRandomArray(size int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	var randomArray = make([]int, size)
	for i := 0; i < size; i++ {
		randomArray[i] = rand.Intn(maxValue)
	}
	return randomArray
}
