package main

import "fmt"

func main() {
	// комплексные: complex64, complex128;
	// вещественные: float32, float64.
	var num32 float32
	var num64 float64

	num32 = -111111111111111111111111111111111111111.02133331313123
	num64 = 1111111111111111111111111111111111111111111111111111111111111111111.1231412

	fmt.Println("float32 -", num32)
	fmt.Println("float64 -", num64)
}
