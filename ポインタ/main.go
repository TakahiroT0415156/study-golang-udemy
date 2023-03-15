package main

import "fmt"

func Double(i int) {
	i = i * 2
}

func Double2(i *int) {
	*i = *i * 2
}

func main() {
	var n int = 100
	fmt.Println(n)

	// var p *int = &n
	// fmt.Println(p)
	// fmt.Println(*p)

	// *p = 300
	// Double(*p)
	Double2(&n)
	fmt.Println(n)

}
