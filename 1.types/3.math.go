package main

import "fmt"

func main() {
	// + — сложение
	// - — вычитание
	// * — умножение
	// / — деление нацело
	// % — остаток от деления
	
	var num1 = 10
	var num2 = 7

	fmt.Println("num1 + num2", num1 + num2)
	fmt.Println("num1 - num2", num1 - num2)
	fmt.Println("num1 * num2", num1 * num2)
	fmt.Println("num1 / num2", num1 / num2)
	fmt.Println("num1 % num2", num1 % num2)

	num1++
	num2--

	fmt.Println("num1++", num1)
	fmt.Println("num2--", num2)

	num1 += num2
	fmt.Println("num1 += num2", num1)
}
