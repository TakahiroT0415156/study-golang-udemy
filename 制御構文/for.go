package main

import "fmt"

// for

func main() {
	i := 0
	for {
		i++
		if i == 1 {
			break
		}
		fmt.Println(("Loop"))
	}

	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	arr1 := [3]int{1, 2, 3}
	for i, v := range arr1 {
		fmt.Println(i, v)
	}
	for _, v := range arr1 {
		fmt.Println(v)
	}

	sl := []string{"python", "php", "go"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
