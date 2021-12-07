package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
		pet  string
	}

	daniel := person{
		"Daniel",
		24,
		"Bongy",
	}
	brenda := person{
		name: "Brenda",
		age:  24,
		pet:  "Queso",
	}
	fmt.Println(daniel.pet)
	fmt.Println(brenda.pet)
}
