package main

import "fmt"

const (
    one = iota + iota + 1
    three
    five
    seven
    nine
    eleven
)

func main() {
	// Result: 1 3 5 7 9 11
    fmt.Println(one, three, five, seven, nine, eleven)
}