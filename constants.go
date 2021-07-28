package main

import (
	"fmt"
	"unsafe"
)

const(
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {
	println(a, b, c)

	const LENGTH int = 10
	const WIDTH int = 2
	var area int

	const a, b,c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Println("面积为 : %d", area)

	println()
	println(a, b, c)

	iotaFunc()
	iotaUsage()
}

func iotaFunc() {
	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Println(a, b, c, "aa")
}

func iotaUsage() {
	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)

	const (
		i1=1<<iota
		j=3<<iota
		k
		l
	)

	fmt.Println("i=",i1)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
}


