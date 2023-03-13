package main

import "fmt"

func main() {
	byteA := []byte{72, 73}
	// 配列を表現できる型
	fmt.Println(byteA)

	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))

	arr := []byte("日本語")
	fmt.Println(arr)
}
