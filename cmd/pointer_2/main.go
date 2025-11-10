package main

import "fmt"

func main() {
	a := 4
	squareAdd(&a)
	fmt.Println(a)
}

// immutability
func squareVal(v int) {
	v *= v
	fmt.Println(&v, v)
}

// efficiency
func squareAdd(v *int) {
	*v *= *v
}
