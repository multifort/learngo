package queue

import (
	"io"
	"os"
)

//通过定义别名对一个已有类型进行扩展，
type Queue []interface{} //interface{}是任何类型的表示

func (q *Queue) Push(value interface{}) {
	*q = append(*q, value)
	os.Open()
}


func (q *Queue) Pop() interface{} {
	value := (*q)[0]
	*q = (*q)[1:]
	return value.(int)//强制返回值为int
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
