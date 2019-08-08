package main

import (
	"fmt"
	"os"
)

type MyType string
type Number int

type Liters float64
type Milliliters float64
type Gallons float64

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func (m MyType) ExportedMethod() {
}

func (m MyType) unexportedMethod() {
}

func (m MyType) WithReturn() int {
	return len(m)
}

func (n *Number) Double() {
	*n *= 2
}

const GALLONS_PER_LITER float64 = 0.264
const LITERS_PER_GALLON float64 = 1 / GALLONS_PER_LITER
const MILLILITERS_PER_GALLON float64 = 1 / GALLONS_PER_LITER * 1000

func (l Liters) ToGallons() Gallons {
	return Gallons(l * Liters(GALLONS_PER_LITER))
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * Milliliters(GALLONS_PER_LITER) / 1000)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * Gallons(LITERS_PER_GALLON))
}

func main() {
	fmt.Println(os.Args)

	value := MyType("a MyType value")
	value.sayHi()

	anotherValue := MyType("another value")
	anotherValue.sayHi()

	x := 3.14
	fmt.Println(x)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", value)

	anotherOne := MyType("MyType value")
	anotherOne.MethodWithParameters(4, true)

	anotherTwo := MyType("MyType value")
	fmt.Println(anotherTwo.WithReturn())

	anotherThree := MyType("MyType value is all the way to 3!")
	fmt.Println(anotherThree.WithReturn())

	number := Number(4)
	fmt.Println("Original:", number)
	number.Double()
	fmt.Println("Doubled:", number)

	stringy := MyType("knotty")
	pointer := &stringy

	stringy.method()
	stringy.pointerMethod()

	pointer.method()
	pointer.pointerMethod()

	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())

	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())

}
