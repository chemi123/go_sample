package main

/*
#include <sample.hpp>
#include <stdlib.h>
#cgo CFLAGS: -I./sample
#cgo LDFLAGS: -lsample -L./ -Wl,-rpath,./
*/
import "C"
import "unsafe"
import "fmt"

func main() {
	a := C.num()
	fmt.Println(a)
	C.hoge()

	str := C.CString("from go string")
	C.showGoString(str)
	C.free(unsafe.Pointer(str))

	fmt.Println(C.GoString(C.fuga()))

	request := C.CString("this is a request")
	response := C.handlerInterface(request)
	fmt.Println(C.GoString(response))
	C.free(unsafe.Pointer(request))
	// C.free(unsafe.Pointer(response))
}
