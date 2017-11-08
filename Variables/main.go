package main

import "fmt"

func main() {
	//declare variable
	var str string
	var a, b int
	var ab float32

	//assign value
	str = "Hello World"
	a = 10
	b = 50
	ab = 2.50

	fmt.Println("str =", str)
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("ab =", ab)

	//declare variable and assign values to variables
	var kota string = "Bandung"
	var c int = 100

	fmt.Println("kota =", kota)
	fmt.Println("c =", c)

	//declare variable with defining its type
	negara := "Indonesia"
	d := 20

	fmt.Println("Negara =", negara)
	fmt.Println("d =", d)

	//define multiple variables
	var (
		nama  string
		email string
		umur  int
	)

	nama = "Alif"
	email = "soraki0990@gmail.com"
	umur = 18

	fmt.Println("Nama =", nama)
	fmt.Println("Email =", email)
	fmt.Println("Umur =", umur)
}
