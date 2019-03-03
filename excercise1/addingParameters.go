package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more parameters.")
		os.Exit(1)
	}
	arguments := os.Args
	var sum int = 0

	for i := 1; i < len(arguments); i++ {
		fmt.Println(arguments[i])
		arg, err := strconv.Atoi(arguments[i])
		if err == nil {
			sum += arg
		} else {
			fmt.Println("one of parameter is not the integer type.")
			os.Exit(1)
		}
	}

	fmt.Println(sum)
}
