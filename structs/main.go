package main

import (
	"fmt"
)

type subscriber struct {
	name   string
	rate   float64
	active bool
}

var subscriber1 struct {
	name   string
	rate   float64
	active bool
}

var subscriber2 struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly Rate:", s.rate)
	fmt.Println("Active?:", s.active)
}

func defaultSubscriber(name string) subscriber {
	var s subscriber

	s.name = name
	s.rate = 5.99
	s.active = true

	return s
}

func main() {
	subscriber1.name = "Steve Ingram"
	subscriber1.rate = 4.99
	subscriber1.active = true

	fmt.Println("Name:", subscriber1.name)
	fmt.Println("Monthly Rate:", subscriber1.rate)
	fmt.Println("Active?:", subscriber1.active)

	fmt.Printf("%#v\n", subscriber1)

	subscriber2.name = "Kerstin Dengl"
	subscriber2.rate = 5.99
	subscriber2.active = false

	fmt.Println("Name:", subscriber2.name)
	fmt.Println("Monthly Rate:", subscriber2.rate)
	fmt.Println("Active?:", subscriber2.active)

	fmt.Printf("%#v\n", subscriber2)

	var subscriber3 subscriber

	subscriber3.name = "Alex Letourneau"
	subscriber3.rate = 0.99
	subscriber3.active = true

	fmt.Println("Name:", subscriber3.name)
	fmt.Println("Monthly Rate:", subscriber3.rate)
	fmt.Println("Active?:", subscriber3.active)

	fmt.Printf("%#v\n", subscriber3)

	subscriber4 := defaultSubscriber("Mackenzie Rudis")
	subscriber4.rate = 2.99
	printInfo(subscriber4)

	subscriber5 := defaultSubscriber("Chris Bojemski")
	printInfo(subscriber5)
}
