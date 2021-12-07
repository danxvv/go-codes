package main

import "fmt"

func main() {
	// var nilMap map[string]int
	// fmt.Println(nilMap == nil)
	// emptyMap := map[string]int{}
	// fmt.Println(emptyMap == nil)
	// teams := map[string][]string{
	// 	"Orcas": {"Daniel", "Juan"},
	// 	"Lions": {"Paco", "Jaime"},
	// }
	// fmt.Println(teams)
	// totalWins := map[string]int{}
	// totalWins["Orcas"] = 2
	// totalWins["Lions"] = 4
	// fmt.Println(totalWins["Orcas"])
	// fmt.Println(totalWins["Lions"])
	// fmt.Println(totalWins["Kittiens"])
	// totalWins["Kittiens"] = 3
	// fmt.Println(totalWins["Kittiens"])
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)

}
