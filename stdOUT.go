package main

import (
	"os"
	"io"
)

func main(){
	myString := ""
	arguments := os.Args;
	if len(arguments) ==1 {
		myString = "PLZ give me one arguments"
	} else {
		myString = arguments[0] + arguments[1];
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}