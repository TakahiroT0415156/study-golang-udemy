package main

import "fmt"

// 並行処理を行うもの

func main() {
	// 受信の書き方
	// var ch2 <-chan int
	// 送信専用
	// var ch3 chan<- int

	ch1 := make(chan int)
	fmt.Println(cap(ch1))

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	ch3 <- 1

	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println(len(ch3))

	i := <-ch3
	fmt.Println(i)

	i2 := <-ch3
	fmt.Println(i2)
}
