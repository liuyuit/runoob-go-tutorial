package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

type Animal interface {
	speak()
}

type Cat struct {
}

func (cat Cat) speak() {
	fmt.Println("cat is speaking")
}

type Dog struct {
}

func (dog Dog) speak() {
	fmt.Println("dog is speaking")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	var animal Animal
	animal = new(Cat)
	animal.speak()
	animal = new(Dog)
	animal.speak()
}
