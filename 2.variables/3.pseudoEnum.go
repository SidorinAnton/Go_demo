package main

import "fmt"

func main() {
	const (
		Black = iota
		Gray
		White
	)

	// счётчик обнуляется
	const (
		Yellow = iota
		Red
		Green = iota // это присваивание не обнулит iota
		Blue
	)

	fmt.Println(Black, Gray, White)       // 0 1 2
	fmt.Println(Yellow, Red, Green, Blue) // 0 1 2 3
}
