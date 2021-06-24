package main

// int plus(int a, int b) {
//     return a + b;
// }
import "C"

func main() {
	println(C.plus(C.int(4), C.int(7)))
}
