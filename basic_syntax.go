package main

import "fmt"

func main() {
	fmt.Println("Google" + "Runoob")

	var age int
	fmt.Printf("age: %d\n", age)

	var stockCode = 123
	var endDate = "2020-12-31"
	var url = "code=%d&date=%s"
	var tarUrl = fmt.Sprintf(url, stockCode, endDate)

	fmt.Println(tarUrl)
}