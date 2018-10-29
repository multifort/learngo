package main

import "fmt"

//循环的语法要求，首先需要明确go语言中不存在while语句。for循环即可以当做循环读取数据，
// 也可以用来做条件控制语句，还可以迭代操作
//for循环完整的语法定义
func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

//for循环省略初始化数据的条件定义；
func looper(num int) {
	for ; num < 10; num ++ {
		fmt.Println(num)
	}
}

//for循环只保留判断条件的初始化数据定义
func loopExp() {
	num := 5
	for num < 10  {
		fmt.Println(num)
		num ++
	}
}

func main() {
	loop()
	looper(8)
	loopExp()
}
