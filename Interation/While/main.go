package main

import "fmt"

func main() {
	//iteration - while
	// go doesn't provide while syntax. we can use for

	a := 1
	for a <= 6 {
		fmt.Println(a)
		a++
	}
}
