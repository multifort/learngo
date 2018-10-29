package main

import (
	"fmt"
)

//struct结构体的格式定义
type treeNode struct {
	value       int
	left, right *treeNode
}

//struct结构体初始化方法4

//struct结构体函数定义
func createNode() *treeNode {
	return new(treeNode)
}

//struct结构体修改值
func (node *treeNode) print() {
	fmt.Println(node.value, " ")
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Node is nil, ingored")
	}
	node.value = value
}

//按照中序遍历树节点
func (node *treeNode) traverse() {

	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	//struct结构体初始化方法1
	var root treeNode
	fmt.Println(root)
	fmt.Println("------")
	//struct结构体初始化方法2
	root.left = &treeNode{value: 3}
	root.setValue(20)
	root.print()
	fmt.Println("------")
	fmt.Println(root.left)
	//struct结构体初始化方法3
	root.right = new(treeNode)
	fmt.Println(root)
	root.right.value = 4
	root.right.print()
	root.right.setValue(10)
	root.right.print()

	root.left.left = createNode()
	root.left.left.setValue(9)
	root.left.left.print()

	fmt.Println("0000000000-------00000000000")
	root.traverse()

}
//struct传值和传指针的区别, 需要明确的是在go语言中没有class的概念，也没有封装和继承的概念，go语言通过面向函数编程，
//通过定义struct中以及struct中定义方法完成对struct的操作，其中struct中没有构造方法。
//只有指针可以修改方法，nil指针也可以调用方法，和其他语言的null不同

//值传递与指针传递的区别
//1.要改变内容必须使用指针接收者
//2.结构过大的时候，考虑使用指针接收者
//3.一致性考虑，如果有指针接收者，最好使用指针接收者