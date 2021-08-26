package main

import "fmt"

var g int
var g1 int = 10

func main() {
	var a, b, c int
	a = 10
	b = 20
	c = a + b

	fmt.Println(a, b, c)
	globalFunc()
}

func globalFunc() {
	var a, b int
	a = 10
	b = 20
	g = a + b

	fmt.Println(a, b, g, g1)

	var g1 = 12
	var aa = 12
	fmt.Println(g1, aa)
}
