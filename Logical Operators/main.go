package main

import "fmt"

func main() {
	var (
		a = 5
		b = 10
	)

	// && is and
	fmt.Println(b > a && a != b)
	// ! is not
	fmt.Println(!(a >= b))
	// || is or
	fmt.Println(a == b || a > b)
}
