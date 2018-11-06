package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	/*for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %c \n", id, n)
	}*/

	for n := range c{
		fmt.Printf("Worker %d received %c \n", id, n)
	}
	/*for {
		fmt.Printf("Worker %d received %c \n", id, <-c)
	}*/
}

func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func bufferChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	time.Sleep(time.Millisecond)
}

func closeChannel() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	time.Sleep(time.Millisecond)
}

func channelDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func main() {
	//channelDemo()
	bufferChannel()
	closeChannel()
}
