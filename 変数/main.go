package main

import "fmt"

// 関数外で暗黙的定義は不可
// i5 := 500

var i5 int = 500

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

// 変数定義
func main() {
	// 明示的な定義
	// var 変数名 型 = 値

	// 数値の定義
	var i int = 100
	fmt.Println((i))

	// 文字列の定義
	var s string = "Hello World"
	fmt.Println(s)

	// 真偽値の定義
	var t, f bool = true, false
	fmt.Println(t, f)

	// まとめて定義することができる
	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	// 変数だけを型定義する
	// 型の初期値が入るようになる
	// intの初期値は0
	var i3 int
	// stringの初期値は空文字が入る
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Golangの勉強中"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	// 暗黙的な定義
	// 変数 := 値
	// 型指定の必要がない
	i4 := 400
	fmt.Println(i4)

	i4 = 500
	fmt.Println(i4)

	// i4 := 800のような再宣言は不可

	// i4 = "Hello"のようなint型にstring型を入れようとするとエラーが起こる

	fmt.Println(i5)

	// 暗黙定義と明示的定義の使い分けは主に関数内か関数外で分ける
	// 基本的には明示的定義を行う方が良い

	// 関数の実行方法
	outer()

	// Golangは定義した変数は使わなければエラーが起こるようになっている
}
