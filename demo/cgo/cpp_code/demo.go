package main

// #cgo darwin CFLAGS: -I$./lib_darwin
// #cgo darwin LDFLAGS: -L./lib_darwin -lfoo -lstdc++
// #cgo linux CFLAGS: -I$./lib_linux
// #cgo linux LDFLAGS: -L./lib_linux -lfoo -lstdc++
// #include "bridge.h"
import "C"

func main() {
	C.call()
}
