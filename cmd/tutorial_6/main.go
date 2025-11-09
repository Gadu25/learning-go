package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
	int
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	// var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}, 10}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name, myEngine.int)

	// non reusable sample
	var myStruct = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Println(myStruct.mpg, myStruct.gallons)

	fmt.Printf("Total miles left in my tank is: %v\n", myEngine.milesLeft())
	canMakeIt(myEngine, 50)

	var myElectricEngine electricEngine = electricEngine{15, 15}

	canMakeIt(myElectricEngine, 50)
}
