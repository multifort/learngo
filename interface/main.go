package main

import (
	"fmt"
	"learngo/interface/retriver"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Multifort interface {
	Do(value int) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	//r = retriver.Retriever{}
	r = &retriver.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	fmt.Printf("%T %v \n", r, r)
	//inspect(r)
	//fmt.Println(download(r))
}
