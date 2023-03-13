package main

import "fmt"

// 型
// 整数型

func main() {
	var i int = 100
	// このint定義は環境依存型の指定
	fmt.Println(i + 50)

	// intにはint8, int16, int32, int64の最小値最大値が違う指定の仕方ができる

	var i2 int64 = 200
	fmt.Println(i2)
	// fmt.Println(i2 + i)
	// int型が違う場合はエラーが起こる
	// fmt.Println(i2 + int64(i))
	// だとエラーが起きないようにすることができる

	fmt.Printf("%T\n", i2)
	// 型を調べる関数

	fmt.Printf("%T\n", int32(i2))
	// int型を64から32などに変更できる
}
