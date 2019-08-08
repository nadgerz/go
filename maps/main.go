package main

import (
	"fmt"
)

func main() {
	var map1 map[string]int
	fmt.Printf("%#v\n", map1)
	map1 = make(map[string]int)
	fmt.Printf("%#v\n", map1)

	map1["gold"] = 1
	map1["silver"] = 2
	map1["bronze"] = 3

	fmt.Println(map1["bronze"])
	fmt.Println(map1["gold"])

	map2 := make(map[string]int)
	fmt.Printf("%#v\n", map2)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"
	fmt.Println(elements["Li"])
	fmt.Println(elements["H"])
	fmt.Printf("%#v\n", elements)

	isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true
	fmt.Println(isPrime[4])
	fmt.Println(isPrime[7])
	fmt.Printf("%#v\n", isPrime)

	ranks := map[string]int{"bronze": 3}
	fmt.Println(map1["bronze"])
	fmt.Printf("%#v\n", ranks)

	stuff := map[string]string{
		"H":  "Hydrogen",
		"Li": "Lithium",
	}
	fmt.Println(stuff["Li"])
	fmt.Printf("%#v\n", stuff)

}
