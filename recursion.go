package main

import "fmt"

func main() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))

	var i1 int
	for i1 = 0; i1 < 10; i1++ {
		fmt.Printf("%d\t", fibonacci(i1))
	}
}

func Factorial(n uint64) (result uint64) {
	if n <= 0 {
		return 1
	}

	result = n * Factorial(n-1)
	return result
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
