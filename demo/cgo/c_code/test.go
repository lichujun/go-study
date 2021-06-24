package main

//#include "test.h"
import "C"

func main() {
	println(C.plus(C.int(1), C.int(2)))
}
