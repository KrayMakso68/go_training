package lang_basics

import "fmt"

// var anons
var a string
var b, c, d string

var e string = "hi"

var f = "hello"

var (
	g bool = true
	h bool = false
)

func short_annons() {
	name := "c"
	fmt.Println(name)
}

var name, age = "Tom", 23
