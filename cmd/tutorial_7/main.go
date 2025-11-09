package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)

	// *p = 10
	p = &i
	*p = 1

	fmt.Printf("\nThe value p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)
}
