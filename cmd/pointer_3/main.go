package main

import "fmt"

type person struct {
	name string
	age  int
}

func initPerson() *person {
	m := person{name: "noname", age: 50}
	fmt.Printf("initPerson ---> %p\n", &m)
	return &m
}

func main() {
	// value of m were saved through garbage collectors:
	// just make sure don't rely on this too much
	// cause overloading garbage collectors affects performance.
	fmt.Printf("main ---> %p\n", initPerson())
}
