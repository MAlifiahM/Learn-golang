package main

import "fmt"

func main() {
	var a, b int

	a = 10
	b = 5

	c := a + b
	fmt.Printf("%d + %d =%d \n", a, b, c)
	d := a - b
	fmt.Printf("%d - %d =%d \n", a, b, d)
	e := float32(a) / float32(b)
	fmt.Printf("%d / %d =%.2f \n", a, b, e)
	f := a * b
	fmt.Printf("%d * %d =%d \n", a, b, f)
}
