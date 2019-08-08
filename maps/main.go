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

}
