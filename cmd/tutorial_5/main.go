package main

import (
	"fmt"
	"strings"
)

func main() {
	// var myString = "résumé"
	var myString = []rune("résumé")
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\nThe length of 'myString' is %v", len(myString))

	// single quote is a rune
	var myRune = 'a'
	fmt.Printf("\nmyRune = %vW", myRune)

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var catStr = ""
	// this create a completely new string anytime, which is inefficient.
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)
	// strings are immutable
	// catStr[0] = 'a'

	// more efficient way use a string package
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	// it only creates a string in the end.
	var catStr2 = strBuilder.String()
	fmt.Printf("\n%v", catStr2)
}
