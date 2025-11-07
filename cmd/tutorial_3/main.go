package main

import "fmt"

func main() {
	// var intArr [3]int32
	// intArr[1] = 123
	// fmt.Println(intArr[1:3])

	// var intArr [3]int32 = [3]int32{1, 2, 3}
	// intArr := [3]int32{1, 2, 3}
	intArr := [...]int32{1, 2, 3, 4}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{5, 6, 7}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 8)
	fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{9, 10}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice3), cap(intSlice3))

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 32, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	var name string = "Sarah"
	var age, ok = myMap2[name]
	if ok {
		fmt.Printf("%v's age is %v\n", name, age)
	} else {
		fmt.Println("Can't find person to the map")
	}

	// delete from map !NO RETURNING VALUE
	// delete(myMap2, "Adam")

	// loop through map
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, age: %v\n", name, age)
	}

	// loop through array/slice
	for i, v := range intArr {
		fmt.Printf("Index %v, Value %v\n", i, v)
	}

	for i, v := range intSlice3 {
		fmt.Printf("Index %v, Value %v\n", i, v)
	}

	// so there's no while loop in go.
	// var i int = 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// or
	// for {
	// 	if i >= 9 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	// or
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
