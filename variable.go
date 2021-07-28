package main

import "fmt"

var x, y int = 12, 13

var (
	xa int = 12
	xb string = "ok"
)

func main() {
	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
	fmt.Println(x, y)
	fmt.Println(xa, xb)

	zeroVariable()
	multiDefine()
}

func zeroVariable() {
	var a string = "Runoob"
	fmt.Println(a)

	var b int
	fmt.Println(b)

	var i int
	var f float64
	var bo bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, bo, s)

	var d = true
	fmt.Println(d)

	varS := "Runoob"
	fmt.Println(varS)
}

func multiDefine() {
	var vname1, vname2, vname3 int
	fmt.Println(vname1, vname2, vname3)

	vname11, vname12, vname13 := 3, "3", true
	fmt.Println(vname11, vname12, vname13)
}