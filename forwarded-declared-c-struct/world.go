// +build ignore

package main

// #cgo CFLAGS: -I${SRCDIR}/cpp
// #cgo LDFLAGS: -L${SRCDIR}/cpp/build -lhello
// #cgo LDFLAGS: -Wl,-rpath=${SRCDIR}/cpp/build
// #include <stdlib.h>
// #include "cpp/hello.hpp"
import "C"

import "fmt"

func main() {

	x := new(C.App)
	y := new(C.App)

	fmt.Printf("x: %p\n", x)
	fmt.Printf("y: %p\n", y)

	a := new(C.App2)
	b := new(C.App2)

	fmt.Printf("a: %p\n", a)
	fmt.Printf("b: %p\n", b)
}
