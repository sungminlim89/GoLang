package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("give a one or more aruments.")
		os.Exit(1)
	}

	arguments := os.Args

	for i := 1; i < len(arguments); i++ {
		if strings.Compare(arguments[i], "STOP") == 0 {
			fmt.Println("normally over.")
			os.Exit(1)
		} else {
			now, err := strconv.ParseFloat(arguments[i], 64)
			if err != nil {
				fmt.Println("it's not a real number.")
				os.Exit(1)
			} else {
				fmt.Print(now)
				fmt.Print(" ")
			}
		}
	}
}
