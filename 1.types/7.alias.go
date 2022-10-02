package main

import "fmt"

func main() {
	type MyString = string // MyString здесь — это псевдоним типа string

	var a string 
	var b MyString
	a = b 

	fmt.Println("a, b", a, b)
}
