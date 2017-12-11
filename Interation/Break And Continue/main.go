package main

import "fmt"

func main() {
	//iteration - for
	var a int
	for a = 0; a < 6; a++ {
		if a == 3 {
			break
		}
		fmt.Println(a)
	}
	for b := 6; b < 11; b++ {
		if b == 7 {
			continue
		}
		fmt.Println(b)
	}
}
