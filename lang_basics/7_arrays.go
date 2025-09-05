package lang_basics

import "fmt"

func arrays() {
	var numbers [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(numbers)

	numbers = [4]int{1, 2}
	fmt.Println(numbers) // [1 2 0 0]

	numbers2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers2)

	numbers3 := [...]int{1, 2, 3} // длина массива 3
	fmt.Println(numbers3)         // [1 2 3]

	numbers[0] = 87
	fmt.Println(numbers[0]) // 87

	multiNumbers := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(multiNumbers)

	multiNumbers[0] = [2]int{66, 77}
	fmt.Println(multiNumbers)

	massLen := len(numbers)
	fmt.Println(massLen)

}
