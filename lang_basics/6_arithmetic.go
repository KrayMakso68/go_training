package lang_basics

import "fmt"

func arithmetic() {
	var a = 4
	var b = 6
	var c = a + b
	fmt.Println(c)

	c = a * b
	fmt.Println(c)

	c = a / b
	fmt.Println(c) //1

	var g float64
	g = 10 / 4.0
	fmt.Println(g) //2.5

	c = 35 % 3 // 35-33=2

	a++
	a--
}
