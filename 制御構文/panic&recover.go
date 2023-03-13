package main

import "fmt"

// 例外処理

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	panic("runtime error")
	// fmt.Println("Start")
}
