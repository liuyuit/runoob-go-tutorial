package main

import "fmt"

func main() {
	var a int = 10

	fmt.Printf("变量的地址: %x\n", &a)
	usage()
	nilPtr()
}

func usage() {
	var a int = 20
	var ip *int
	ip = &a

	fmt.Println(&a)
	fmt.Println(*ip)
	fmt.Println(ip)
	fmt.Println(&ip)
}

func nilPtr() {
	var ptr *int
	fmt.Println(ptr)

	if ptr == nil {
		fmt.Println("this ptr is nil")
	}
}
