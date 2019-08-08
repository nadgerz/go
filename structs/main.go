package main

import (
	"fmt"
)

var subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	subscriber.name = "Steve Ingram"
	subscriber.rate = 4.99
	subscriber.active = true

	fmt.Println("Name:", subscriber.name)
	fmt.Println("Monthly Rate:", subscriber.rate)
	fmt.Println("Active?:", subscriber.active)

	fmt.Printf("%#v\n", subscriber)
}
