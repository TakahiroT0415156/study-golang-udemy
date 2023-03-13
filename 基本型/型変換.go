package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	fl64 := float64(i)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)

	i2 := int(fl64)
	fmt.Printf("i2 = %T\n", i2)

	var str string = "100"
	fmt.Printf("s = %T\n", str)

	strToi, _ := strconv.Atoi(str)
	// _にすることで2つ目に関しては消すとができる
	fmt.Println(strToi)
	fmt.Printf("strToi = %T\n", strToi)
}
