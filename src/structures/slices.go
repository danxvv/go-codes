package main

import "fmt"

func main() {
	// var x = []int{1, 2, 3, 4} //[...] array, []slice
	// var x []int // crea un nil
	// var x []int // crea un nil
	// x = append(x, 2)
	// x = append(x, 2)
	// var x = []int{1, 2, 3, 4}
	// x = append(x, 5, 6, 7, 8, 9, 10)
	// y := []int{30, 40, 50}
	// x = append(x, y...)
	//make nos hace una slice con capacidad especifica, asi se evita el nil value
	x := make([]int, 0, 10)
	x = append(x, 2, 3)
	fmt.Println(len(x), cap(x))
}
