package main

import "fmt"

func main() {
	const (
		i         = 10
		f         = 1.5
		i64 int64 = 88
	)

	var v = 45

	fmt.Println(i + f)
	fmt.Println(i + i64)
	// fmt.Println(f + i64) - Ошибка
	// fmt.Println(i64 + v) - Ошибка 
	fmt.Println(i + v)
	// fmt.Println(f + v) - Ошибка 
}
