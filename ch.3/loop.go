package main

import (
	"fmt"
)

func loop1() {
	for i := 0; i < 100; i++ {
		if i % 20 == 0 {
			continue ;
		}

		if i == 95 {
			break;
		}
		fmt.Print(i, " ");
	}
	fmt.Println()
}

func loop2() {
	for i := 10; i >= 0; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func loop3() {
	i := 0
	for ok := true; ok; {
		if i > 10 {
			ok = false;
		}
		fmt.Print(i, " ")
		i++;
	}
	fmt.Println();
}

func loop4() {
	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index: ", i, " value: ", value)
	}
}

func main () {
	loop1()
	loop2()
	loop3()
	loop4()
}
