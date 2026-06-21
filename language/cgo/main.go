package main

/*
#include <stdlib.h>
#include <stdio.h>
*/
import "C"

import "unsafe"

func main() {
	message := C.CString("Hello, cgo")
	defer C.free(unsafe.Pointer(message))
	C.puts(message)
}
