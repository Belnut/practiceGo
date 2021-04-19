package main

// #include <stdio.h>
// void callC () {
// 	printf("Calling C code!\n");
// }
// int returnInt()
//{
//	return 10;
//}
import "C"

import (
	"fmt"
)

func main(){
	fmt.Println("A Go statement!")
	C.callC()
	val := C.returnInt()
	fmt.Println("Another Go state", val)
}