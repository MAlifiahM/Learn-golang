package main

import "fmt"

func main() {
	var (
		a = 5
		b = 10
	)

	if a > b || a-b < a {
		fmt.Println("conditional--> a > b || a - b < a")
	} else {
		fmt.Println("..another")
	}
}

// decision else must like that or you get an error
