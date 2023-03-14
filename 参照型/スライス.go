package main

import "fmt"

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	sl3 := []string{"a", "b"}
	fmt.Println(sl3)

	sl4 := make([]int, 5)
	fmt.Println(sl4)
	fmt.Println(sl4[2:4]) // 2~4番目の配列の値を取り出すことができる

	fmt.Println(sl4[:2])         // 2番目の手前までの配列の値を取り出すことができる
	fmt.Println(sl4[2:])         // 2番目の最後までの配列の値を取り出すことができる
	fmt.Println(sl4[:])          // 全ての配列の値を取り出すことができる
	fmt.Println(sl4[len(sl4)-1]) // 配列の最後の値を取り出すことができる

	sl2 = append(sl2, 300)
	fmt.Println(sl2)

	sl5 := make([]int, 5)
	fmt.Println(sl5)
	fmt.Println(len(sl5))
	fmt.Println(cap(sl5))

	sl6 := make([]int, 5, 10)
	fmt.Println(sl6)
	fmt.Println(len(sl6))
	fmt.Println(cap(sl6))

	sl7 := []int{100, 200}
	sl8 := sl7
	sl8[0] = 1000
	fmt.Println(sl7) // sl7まで配列のデータが入れ替わってしまう

	sl9 := []int{1, 2, 3, 4, 5}
	sl10 := make([]int, 5, 10)
	n := copy(sl10, sl9)
	fmt.Println(n, sl10)
	sl10[0] = 1000
	fmt.Println(sl10, sl9)

	sl11 := []string{"A", "B", "C"}
	fmt.Println(sl11)

	for i, v := range sl11 {
		fmt.Println(i, v)
	}

	for i := 0; i < len(sl11); i++ {
		fmt.Println(sl11[i])
	}

	fmt.Println(Sum(1, 2, 3))
	sl13 := []int{1, 2, 3}
	fmt.Println(Sum(sl13...))
}
