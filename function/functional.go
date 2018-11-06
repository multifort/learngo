package main

import (
	"fmt"
)

func Add() func(value int) int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum
	}
}

func fabonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	a := Add()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
	b := fabonacci()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}
