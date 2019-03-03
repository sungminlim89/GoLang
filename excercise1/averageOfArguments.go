package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("give a 1 or more floats.")
		os.Exit(1)
	}

	arguments := os.Args
	var sum float64 = 0

	for i := 1; i < len(os.Args); i++ {
		now, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			fmt.Println("one of arguments is not a float type.")
			os.Exit(1)
		} else {
			sum += now
		}
	}

	avg := sum / float64((len(os.Args) - 1))
	fmt.Print("avg: ")
	fmt.Println(avg)
}
