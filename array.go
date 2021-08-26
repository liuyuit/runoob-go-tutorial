package main

import "fmt"

func main() {
	var i, j int = 1, 1
	var n [10]int

	for i < 10 {
		n[i] = i + 100
		i++
	}

	for j < 10 {
		fmt.Println(n[j])
		j++
	}

	ini()
}

func ini() {
	var balance = [3]float32{11.22, 12.12, 10.01}

	//var balanceInt = [...]int{1,2,3}

	length := len(balance)

	for i := 0; i < length; i++ {
		fmt.Println(balance[i])
	}

	var balance1 = [3]float32{0: 1.3, 2: 1.33}

	for k := 0; k < 3; k++ {
		fmt.Println(balance1[k])
	}
}
