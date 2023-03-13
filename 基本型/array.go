package main

import "fmt"

// 配列型

func main() {
	// 他の言語と違って配列の直接の操作はできない
	var arr1 [3]int
	// []の中は要素数を表している
	fmt.Println(arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"a", "b"}
	// 自動で要素数を判断してくれる
	fmt.Println(arr4)
}
