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

	emptyMap := map[string]float64{}
	fmt.Printf("%#v\n", emptyMap)

	counters := make(map[string]int)
	counters["a"]++
	counters["a"]++
	counters["c"]++
	fmt.Println(counters["a"], counters["b"], counters["c"])
	fmt.Printf("%#v\n", counters)

	counters2 := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	value, ok = counters2["a"]
	fmt.Println(value, ok)
	value, ok = counters2["b"]
	fmt.Println(value, ok)
	value, ok = counters2["c"]
	fmt.Println(value, ok)
	fmt.Printf("%#v\n", counters2)

	var rank int
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	fmt.Printf("%#v\n", ranks)

}
