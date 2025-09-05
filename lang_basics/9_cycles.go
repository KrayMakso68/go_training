package lang_basics

import "fmt"

func cycles() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d\t", i)
	}

	i := 1
	for ; i < 100; i++ {
		fmt.Printf("%d\t", i)
	}

	i = 10
	for i <= 20 {
		fmt.Printf("%d\t", i)
		i++
	}

	str := "Hello"
	for index, value := range str {
		fmt.Println("Index:", index, " Value:", value)
	}

	for _, value := range str {
		fmt.Println("Value:", value)
	}

	var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
	var sum = 0

	for _, value := range numbers {
		if value < 0 {
			continue // переходим к следующей итерации
		}
		sum += value
	}
	fmt.Println("Sum:", sum) // Sum: 27

	for _, value := range numbers {
		if value > 4 {
			break // если число больше 4 выходим из цикла
		}
		sum += value
	}
	fmt.Println("Sum:", sum) // Sum: 10

OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {

			fmt.Printf("i = %d, j = %d\n", i, j)

			if i == 2 && j == 2 {

				fmt.Println("Выход из внешнего цикла...")

				break OuterLoop // выходим из внешнего цикла
			}
		}
	}

	fmt.Println("Цикл завершен...")
}
