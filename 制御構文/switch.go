package main

import "fmt"

// switch

func anything(a interface{}) {
	// fmt.Println(a)
	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't Know")
	}
}

func main() {
	// 型switch
	anything("aaa")
	anything(1)

	var x interface{} = 3
	i := x.(int)
	fmt.Println((i + 2))

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't Know")
	}

	// 式switch
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1or2")
	case 3, 4:
		fmt.Println("3or4")
	default:
		fmt.Println("I don't Know")
	}

	switch n1 := 5; n1 {
	case 1, 2:
		fmt.Println("1or2")
	case 3, 4:
		fmt.Println("3or4")
	default:
		fmt.Println("I don't Know")
	}

	n2 := 6
	switch {
	case n2 > 0 && n2 < 4:
		fmt.Println("0<n<4")
	case n2 > 3 && n2 < 7:
		fmt.Println("3<n<7")
	}
}
