package lang_basics

import "fmt"

func types() {
	var a int8 = -1
	var b uint8 = 2
	var c byte = 3 // byte - синоним типа uint8
	var d int16 = -4
	var f uint16 = 5
	var g int32 = -6
	var h rune = -7 // rune - синоним типа int32
	var j uint32 = 8
	var k int64 = -9
	var l uint64 = 10
	var m int = 102
	var n uint = 105

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	fmt.Println("f: ", f)
	fmt.Println("g: ", g)
	fmt.Println("h: ", h)
	fmt.Println("j: ", j)
	fmt.Println("k: ", k)
	fmt.Println("l: ", l)
	fmt.Println("m: ", m)
	fmt.Println("n: ", n)

	var o float32 = 18
	var p float32 = 4.5
	var q float64 = 0.23
	var pi float64 = 3.14
	var e float64 = 2.7

	fmt.Println("f: ", o)
	fmt.Println("g: ", p)
	fmt.Println("d: ", q)
	fmt.Println("pi: ", pi)
	fmt.Println("e: ", e)

	var r complex64 = 1 + 2i
	var s complex128 = 4 + 3i

	fmt.Println("r: ", r)
	fmt.Println("s: ", s)

	var isAlive bool = true
	var isEnabled bool = false
	fmt.Println("isAlive: ", isAlive)
	fmt.Println("isEnabled: ", isEnabled)

	var name string = "Hello Go"
	fmt.Println("name: ", name)
	/*
		\n: переход на новую строку
		\r: возврат каретки
		\t: табуляция
		\": двойная кавычка внутри строк
		\\: обратный слеш

		%T: тип переменной в выводе Println
	*/
}
