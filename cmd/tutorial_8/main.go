package main

import "fmt"

// the main exception of this copy behavior of non-pointer variable
// is when working with slices
func main() {
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice

	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	// because under the hood slices contain pointers to an underlying array
}
