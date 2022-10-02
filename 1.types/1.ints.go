package main

import "fmt"

func main() {
	// знаковые: int, int8, int16, int32, int64;
	// беззнаковые: uint, uint8, uint16, uint32, uint64.
	var num int
	var num8 int8
	var num16 int16
	var num32 int32
	var num64 int64
	var unum uint
	var unum8 uint8
	var unum16 uint16
	var unum32 uint32
	var unum64 uint64

	num = -1
	num8 = -3
	num16 = -123
	num32 = 341342141
	num64 = 4130492941294992903
	unum = 1
	unum8 = 111
	unum16 = 11111
	unum32 = 1111111111
	unum64 = 11111111111111111111

	fmt.Println("int -", num)
	fmt.Println("int8 -", num8)
	fmt.Println("int16 -", num16)
	fmt.Println("int32 -", num32)
	fmt.Println("int64 -", num64)
	fmt.Println("uint -", unum)
	fmt.Println("uint8 -", unum8)
	fmt.Println("uint16 -", unum16)
	fmt.Println("uint32 -", unum32)
	fmt.Println("uint64 -", unum64)
}
