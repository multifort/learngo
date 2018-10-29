package main

import (
	"fmt"
)

//定义单个常量
const file = "aa.txt"

//定义一组常量
const (
	aaa = 1
	bbb = "test"
	ccc = true
)

//定义自增性常量
const (
	ddd = iota
	fff
)

//利用常量定义枚举
func enum() {
	const (
		car   = "car"
		bus   = "bus"
		plane = "plane"
	)
	fmt.Println(car, bus, plane)
}

//b,kb,mb,gb,tb,pb的枚举类型
func memEunm() {
	const (
		b  = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println(file)

	fmt.Println(aaa)
	fmt.Println(bbb)
	fmt.Println(ccc)
	fmt.Println(ddd)
	fmt.Println(fff)

	enum()
	memEunm()

}

//go语言没有枚举类型，通过const常量组可以定义一组枚举类型，也可以通过iota自动增长生成一组枚举类型