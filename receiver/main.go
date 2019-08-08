package main

import (
	"fmt"
	"os"
)

type MyType string
type Number int

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
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

}
