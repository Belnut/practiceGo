package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	if len(os. Args) == 1 {
		fmt.Println("PLZ give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64) // 첫번째 값만 가져올것, 나머지는 무시
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i:= 2 ; i< len(arguments); i++ {
		n,_:= strconv.ParseFloat(arguments[i], 64)

		if n < min {
			min = n
		}

		if n> max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}