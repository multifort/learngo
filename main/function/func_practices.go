package function

import "fmt"

//go语言是函数式编程，函数的定义规则，遵循func 函数名称（var 1，2 type） 返回值（可以有多个返回值，包括返回类型）{函数体组成}
//go语言函数分为内置函数和可供外界调用的函数，通过函数名称首字母的大小写来区分

//定义返回值为单个的函数
func test() int {
	return 0
}

//定义返回值为多个的函数
func exchange(a, b int) (int, int) {
	return b, a
}

//定义接受参数为单个的函数
func getOne(i int) int {
	return i
}

//定义接受参数为多个的函数
func multiPramaters(a, b int) {
	fmt.Println(a, b)
}

//定义接受参数为不确定个数的函数
func num(num ... int) {
	for i := range num {
		fmt.Println(i)
	}
}

//定义接受参数为函数的函数
func funcTest(test func()) {
	fmt.Println("test")
}

func test11() ([]int, int) {
	return make([]int, 20), 0
}

func main() {
	test()
	fmt.Println(exchange(3, 2))

	num(1, 2, 3, 4)
	multiPramaters(10, 20)
	funcTest(main)
}
