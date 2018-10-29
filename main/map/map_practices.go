package _map

import (
	"fmt"
)

func mapDef1() {
	var m1 map[string]int
	fmt.Println(m1)
}

func mapDef2() map[string]string {
	m1 := map[string]string{
		"name": "lining",
		"age":  "30",
		"sex":  "male",
	}
	fmt.Println(m1)
	return m1
}

func mapDef3() {
	m1 := make(map[string]int)
	fmt.Println(m1)
}

func modifyMap() {
	m1 := make(map[string]int)
	m1["name"] = 20
	fmt.Println(m1)
	m1["name"] = 30
	fmt.Println(m1)
}

func traversingMap() {
	m1 := mapDef2()
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	for _, v := range m1 {
		fmt.Println(v)
	}

	for k, _ := range m1 {
		fmt.Println(k)
	}
}
func isKeyExised()  {
	m1 := mapDef2()
	name := m1["name"]
	fmt.Println(name)
	if test, ok := m1["test"]; ok{
		fmt.Println(test)
	} else {
		fmt.Println("test is not existed")
	}
}
func delMapKey()  {
	m1 := mapDef2()
	fmt.Println(m1)
	delete(m1, "name")
	fmt.Println(m1)
}

func main() {
	//map的定义格式1-通过 var声明map的定义格式map[keyType]valueType
	mapDef1()
	//map的定义格式2-通过 :=定义map有go语言去做类型的匹配，并且赋初始化值
	mapDef2()
	//map的定义格式3-通过make去定义map,make与var声明的不赋初始化值的区别在于make是一个empty map，而var是一个nil
	mapDef3()
	//map的修改某个key的值
	modifyMap()
	//map的遍历操作
	traversingMap()
	//判断map中的某个key值是否存在
	isKeyExised()
	//map中key的删除操作
	delMapKey()
	//需要注意的是：1.map是基于hash的，因此map中的数据是无序的；
	// 2.map中的key值可以为除了slice，map以及func以外的其他任何类型；
	// 3.定义后有初始化值（Zero value）
}
