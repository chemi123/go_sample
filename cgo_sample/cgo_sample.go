package main

/*
#include <sample.hpp>
#cgo CFLAGS: -I./sample
#cgo LDFLAGS: -lsample -L./ -Wl,-rpath,./
*/
import "C"
import "fmt"

func main() {
        a := C.num()
        fmt.Println(a)
        C.hoge()

	str := C.CString("fuga")
	C.fuga(str)
}

