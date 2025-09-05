package lang_basics

import "fmt"

func consts() {
	const pi float64 = 3.14159265

	var _ int = 7
	// const k = _      // ! Ошибка: m - переменная

	const (
		a = 1
		b
		c
		d = 3
		f
	)
	fmt.Println(a, b, c, d, f) // 1, 1, 1, 3, 3
}

func func_iota() {
	const // iota сбрасывается в 0
	(
		C0 = iota // здесь iota равно 0, увеличивается с каждой строкой
		C1        // увеличение на 1, iota равна 1
		C2 = iota // iota равна 2
	)

	fmt.Println("C0:", C0) // C0: 0
	fmt.Println("C1:", C1) // C1: 1
	fmt.Println("C2:", C2) // C2: 2

	const //  iota сбрасывается в 0
	(
		C3 = iota // С3 = 0
	)

	fmt.Println("C3:", C3) // C3: 0
}
