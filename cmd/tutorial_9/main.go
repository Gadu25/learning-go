package main

import "fmt"

// pointers and functions

// this is inefficient because this creates a copy of a new array.
// the way to improve this is to use pointers of the array,
// and modify the first array using only the pointer address
func main() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square(thing1)
	// to improve
	// var result [5]float64 = square(&thing1)

	fmt.Printf("\nThe result is: %v", result)
}

func square(thing2 [5]float64) [5]float64 {
	// to improve
	// func square(thing2 *[5]float64) [5]float64 {

	fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2)

	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}
