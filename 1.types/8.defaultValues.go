package main

import "fmt"

func main() {
	// bool — false
	// числа — 0
	// string — пустая строка длиной 0
	// ссылки — nil или пустой указатель

	var text string
	var boolVal bool
	var number int

	fmt.Println("Default value for string -", text)
	fmt.Println("Default value for bool- ", boolVal)
	fmt.Println("Default value for int -", number)
}
