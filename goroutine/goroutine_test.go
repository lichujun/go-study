package goroutine

import (
	"fmt"
	"testing"
)

func TestGoroutine(t *testing.T) {
	a := 10
	ch := make(chan int)
	go func() {
		for i := 0; i < a; i++ {
			ch <- i
		}
	}()
	for i := 0; i < a; i++ {
		fmt.Println(<-ch)
	}
}
