package main

import "fmt"

// 定数
// 定数は基本的に他のファイルでも利用するため、グローバルな場所で定義をしていく

const Pi = 3.14

// 頭文字を大文字にすると他のファイルでも呼び出すことができるようになる

const (
	URL      = "http://xxx.com"
	SiteName = "test"
)

const (
	A = 1
	B
	C = "A"
	D
)

func main() {
	fmt.Println(Pi)
	fmt.Println(URL, SiteName)
	fmt.Println(A, B, C, D)
}
