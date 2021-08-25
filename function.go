package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("google", "runoob")
	fmt.Println(a, b)

	x := 1
	y := 3
	result := max(x, y)
	fmt.Printf("max : %d", result)
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
