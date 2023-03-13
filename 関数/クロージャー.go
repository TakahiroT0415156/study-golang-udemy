package main

import "fmt"

// クロージャー

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("My"))     // Hello
	fmt.Println(f("Name"))   // My
	fmt.Println(f("Is"))     // Name
	fmt.Println(f("Golang")) // Is
	fmt.Println(f(""))       // Golang
}
