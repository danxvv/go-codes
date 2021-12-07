package main

import (
	"fmt"
)

func main() {
	// n := rand.Intn(10)
	// fmt.Println(n)
	// if n == 0 {
	// 	fmt.Println("Muy bajo", n)
	// } else if n == 5 {
	// 	fmt.Println("Muy alto", n)
	// } else {
	// 	fmt.Println("Ok", n)
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// i := 1
	// for i < 1000 {
	// 	i *= 2
	// 	fmt.Println(i)
	// }
	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 && i%5 == 0 {
	// 		fmt.Println("FizzBuzz")
	// 		continue
	// 	}
	// 	if i%3 == 0 {
	// 		fmt.Println("Fizz")
	// 		continue
	// 	}
	// 	if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	evenValues := []int{2, 4, 6, 8, 12, 14, 16, 20, 24, 28}
	for _, v := range evenValues {
		fmt.Println(v)
	}

	uniqueValues := map[string]bool{"Fred": true, "Juan": false, "Paco": true}
	for k := range uniqueValues {
		fmt.Println(k)
	}
}
