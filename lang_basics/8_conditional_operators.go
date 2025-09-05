package lang_basics

import "fmt"

func operators() {
	a := 6
	b := 7
	if a < b {
		fmt.Println("a меньше b")
	} else if a > b {
		fmt.Println("a больше b")
	} else {
		fmt.Println("a равно b")
	}

	a = 8

	switch a + 2 {
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	}

	switch a {
	case 3:
		fmt.Println("3")
	case 4, 5, 6, 7:
		fmt.Println("4")
	default:
		fmt.Println("default")
	}

	switch a + 2 {
	case 3:
		fmt.Println("a=3")
		fallthrough
	case 4:
		fmt.Println("a=4")
	}
	// вывод: a=3 a=4

}
