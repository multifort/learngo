package main

import "fmt"

//if语法要求，if标志开始后，条件表达式不需要跟括号，同时if代码块允许定义一个变量，但是变量范围只属于当前代码块
//if条件使用方法（初始化值和条件在同一个if代码块内）
func get()  {
	if n := 40; n <20 {
		fmt.Println(n)
	} else {
		fmt.Println("test")
	}
}


//if条件使用方法（初始化值和条件不在同一个代码块内）
func compat(j int)  {
	if j < 10 {
		fmt.Println(10)
	} else {
		fmt.Println(j)
	}
}

//switch语句的条件可以不存在，不存在时，默认为true，要求switch的条件以及case中的表达式类型必须一致
//switch带条件语句的定义（switch后，跟操作条件等）
func opreator(a,b int, op string) int  {
	switch op {
	case "+":
		return a + b
	case "-":
		return a-b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("test")
	}
}

//swtich不带条件语句的定义（switch后，不带任何限制条件）
//switch中case可以接受函数作为表达式的定义
func grade(score int) string  {
	switch {
	case score < 60:
		return "F"
	case score < 80 :
		return "C"
	case score < 90:
		return "B"
	case score <= 100:
		return "A"
	default:
		panic("test")
	}
}

func main() {
	get()
	compat(2)
	compat(70)

	fmt.Println(opreator(2,3, "-"))
	fmt.Println(grade(10))
}
