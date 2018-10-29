package array

import (
	"fmt"
)

//数组遍历-打印数组下标和数组值
func printArrayKeyAndValue(arr [3]int) {
	for i := range arr {
		fmt.Println(arr[i])
	}

	for i, k := range arr {
		fmt.Println(i, k)
	}
	//go语言中_代表可以忽略的值,range可以用来获取数组或者map或者list的下标值以及具体数值
	for _, k := range arr {
		fmt.Println(k)
	}
}

func main() {
	//数组的定义格式1
	var arr1 [3]int
	//数组的定义格式2
	arr2 := [3]int{1, 2, 3}
	//数组的定义格式3
	arr3 := [...]int{1, 2, 3, 4, 5}
	//定义二维数组
	var graid [4][5]int

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(graid)

	printArrayKeyAndValue(arr2)
}

//需要注意的是，go语言中的数组传递，也包括所有的参数传递都是数值拷贝，即在函数内部改变值，不会影响函数外部值的变化。
//go语言中的数组是一种值类型，和C语言不同的是指针就是指针，例如数组名称在go语言中代表数据类型，不会类似C语言那样代表数组的起始地址