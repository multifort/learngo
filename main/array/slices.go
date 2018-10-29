package array

import (
	"fmt"
)

func createSlice1() {
	var arr [10]int
	for i := range arr {
		arr[i] = i
	}
	s1 := arr[2:5]
	s2 := arr[:5]
	s3 := arr[2:]
	s4 := arr[:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}

func createSlice2() {
	var s []int
	fmt.Println(s)
	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
	}
	s2 := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Println(s2)
}

func createSlice3() {
	var s1 = make([]int, 15)
	var s2 = make([]int, 15, 32)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
}

func viewSliceAndArray() {
	var arr [10]int
	for i := range arr {
		arr[i] = i
	}
	s1 := arr[2:10]
	s2 := s1[:2]
	s3 := s1[7:]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	s2 = append(s2, 10)
	s2 = append(s2, 20)
	fmt.Println(s2)
	fmt.Println(arr)

	s3 = append(s3, 30)
	fmt.Println(s3)
	s3 = append(s3, 40, 40)
	fmt.Println(s3)
	fmt.Println(arr)
}

func showSliceStruct() {
	var s1 = make([]int, 10, 20)
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}

func appendSlice() {
	s1 := make([]int, 10)
	fmt.Println(s1)
	s1 = append(s1, 10, 20, 30, 40)
	fmt.Println(s1)

	s1[0] = 90
	fmt.Println(s1)
}

func delSlice()  {
	var s1= []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(s1)
	s1 = append(s1[:2], s1[3:]...)
	fmt.Println(s1)
}

func combineSlice()  {
	s1 := []int{1,2,3}
	fmt.Println(s1)
	s2 := []int{4,5,6}
	fmt.Println(s2)
	s3 := append(s1, s2...)
	fmt.Println(s3)
}

func delSliceTop()  {
	s1 := make([]int, 10)
	s1[0] = 10
	fmt.Println(s1)
	fmt.Println(s1[0])
	s1 = s1[1:]
	fmt.Println(s1)
}

func delSliceTail()  {
	s1 := make([]int, 10)
	s1[len(s1)-1] = 10
	fmt.Println(s1)
	s1 = s1[:len(s1)-1]
	fmt.Println(s1)
}


func main() {
	//slice的格式定义-方法1-通过获取数组的视图，得到一个slice，可以从头获取，获取中间部分，从某个起始位置开始，以及获取全部数据
	createSlice1()
	//slice的格式定义-方法2-通过不指定数组的大小，直接创建slice,并且给slice赋值
	createSlice2()
	//slice的格式定义-方法3-通过make函数去创建slice，可以指定slice的长度，最大容纳长度等参数
	createSlice3()
	/**slice与array的关系，slice是array的一个视图，当slice中超出array本身的容量以后，
	go语言系统会自动分配一个新的系统空间，并把原来的array数值copy一份，同时在此空间进行操作*/
	viewSliceAndArray()
	//slice的存储结构以及示例组成，包括ptr，len以及cap，分别代表起始位置，当前显示长度以及最大容纳长度
	showSliceStruct()
	//slice中元素增加-通过append方法给slice添加元素,从末尾追加元素,通过指定下标，可以给slice赋值
	appendSlice()
	//slice的元素删除-通过利用append和slice的切分进行元素删除操作,并把剩余元素追加到第一个slice的末尾
	delSlice()
	//两个slice的组合
	combineSlice()
	//删除slice的头元素
	delSliceTop()
	//删除slice的末尾元素
	delSliceTail()
}
