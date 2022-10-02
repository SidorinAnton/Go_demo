package main

import "fmt"

func main() {
	// FizzBuzz 1 -> 100
	// x / 3 -> fizz
	// x / 5 -> buzz
	// x / 3 && x / 5 -> fizzbuzz

	for i := 1; i < 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizzbuzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}
