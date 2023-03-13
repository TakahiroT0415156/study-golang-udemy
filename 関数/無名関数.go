package main

import "fmt"

// 無名関数

func main() {
	f := func(x, y int) int {
		return x + y
	}

	i := f(1, 2)
	fmt.Println(i)

	f2 := func(x, y int) int {
		return x + y
	}(1, 4)
	fmt.Println(f2)
}
