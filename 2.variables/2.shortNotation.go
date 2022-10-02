package main

import "fmt"

func main() {
	num1 := 1
	num2 := 2

	fmt.Println(num1, num2)

	// num1 := 10  -- Нельзя, т.к. эта переменная уже есть
	num1 = 10 // Можно

	fmt.Println(num1)
}
