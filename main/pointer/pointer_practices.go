package pointer

import "fmt"

//传值与传指针的区别，需要明确的是Go语言中的指针只有指向功能，不存在想C语言那样可以对指针进行运算功能。
//传值是在调用过程中对值进行copy，在函数体内对copy的值进行计算，并不会影响调用者内值的变化，
//传指针是指把数据在内存中的地址copy传递到函数内部，函数通过指针和调用者同时指向同一个数据，并且对数据做修改。

//指针定义
var m *int
var str = "test"
var s *string = &str

//获取数据指针
func point(i int) {
	fmt.Println(&i)
}

//函数传递指针，不传递具体值

func pointer(i, j *int) {
	fmt.Println(*j, *i)
}

func main() {
	*s = "aaaa"
	fmt.Println(*s)
	fmt.Println(str)

	point(10)
	//fmt.Println(str)
	//fmt.Println(*s)

	//point(10)
	i, j := 10, 20
	pointer(&i, &j)

	const uuuu  = "test"
	fmt.Println()
}

//可以直接定义指针，但是在指针需要获取地址的时候，必须先定义变量或者常量，然后对变量和常量进行取址操作,
//另外常量是不允许取地址的
//go语言中的指针不存在可操作的行为，只有指向功能，另外go语言中的值在参数之间的传递是传值而不是传址的，
//所以给函数传值并不能改变调用函数中的数据