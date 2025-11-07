package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Howiee")
	numerator := 10
	denominator := 2
	var divRes, divRem, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	} else if divRem == 0 {
		fmt.Printf("The result of the division is: %v", divRes)
	} else {
		fmt.Printf("The result of the division is: %v with remainder of: %v", divRes, divRem)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
	res := numerator / denominator
	var remainder int = numerator % denominator
	return res, remainder, err
}
