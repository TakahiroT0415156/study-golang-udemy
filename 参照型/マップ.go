package main

import "fmt"

// keyとvalueで配列を作る

func main() {
	var m = map[string]int{"a": 100, "b": 200}
	// map[keyの型]valueの型書く
	fmt.Println(m)

	m2 := map[string]int{"a": 100, "b": 200}
	fmt.Println(m2)

	m3 := make(map[int]string)
	fmt.Println(m3)

	m3[0] = "JAPAN"
	m3[1] = "KOREA"
	fmt.Println(m3)
	s, ok := m3[1]
	// 第2引数には値を取れたかどうかのtrue falseが返ってくる
	fmt.Println(s, ok)
}
