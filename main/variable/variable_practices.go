package main

import "fmt"

//定义全局变量方法1（单个）
var aa int = 10

//定义全局变量方法2（多个）
var bb, cc int = 20, 30
var (
	dd string = "test"
	e  int    = 50
)

//定义全局变量方法3（多个）,省略变量类型有程序自己去推算
var (
	a, b, c = 10, "test", true
)

//定义局部变量方法1（全称）
func variableDef() {
	var o int = 0
	var p string = "1232"
	var k bool = true
	fmt.Println(o, p, k)

}

//定义局部变量方法2（简单写法）
func variableDef2() {
	o := 0
	p := "aaa.txt"
	k := true
	fmt.Println(o, p, k)
}

//定义局部变量方法3（多个）
func variableDef3() {
	o, p, k := 1, "bbb.txt", false
	fmt.Println(o, p, k)
}

//打印局部变量的默认值
func variableZeroValue() {
	var o int
	var p string
	var k bool
	fmt.Printf("%s\n", p)
	fmt.Println(o, k)
}

//打印局部变量的初始化值
func variableInitionalValue() {
	o, p, k := 1, "bbb.txt", false
	fmt.Println(o, p, k)
}

//同时输出不同类型的局部变量
func variableTypeValue() {
	o, k := "test.txt", true
	fmt.Println(o, k)
}

//测试定义变量的各个方法
func main() {
	fmt.Println(aa, bb, cc, e, a, b, c)

	variableDef()
	variableDef2()
	variableDef3()
	variableZeroValue()
	variableInitionalValue()
	variableTypeValue()

}

//全局变量和局部变量以及可以省略的类型定义，在全局中不能省略var， 在局部变量中，能够通过 :=用来替换var
