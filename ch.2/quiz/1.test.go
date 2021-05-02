package main

// #cgo LDFLAGS: ${SRCDIR}/1.test.a
// #include <1.test.h>
import "C"

import (
    "fmt"
)

func main() {
    fmt.Println("5 + 10 = ", C.add_num(5, 10))
}
