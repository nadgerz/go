package main

import (
	"fmt"
)

func main() {
	var map1 map[string]int
	fmt.Printf("%#v\n", map1)
	map1 = make(map[string]int)
	fmt.Printf("%#v\n", map1)

	map2 := make(map[string]int)
	fmt.Printf("%#v\n", map2)

}
