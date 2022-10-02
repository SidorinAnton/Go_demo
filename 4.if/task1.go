package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 1946 — 1964: «Привет, бумер!».
	// 1965 — 1980: «Привет, представитель X!».
	// 1981 — 1996: «Привет, миллениал!».
	// 1997 — 2012: «Привет, зумер!».
	// 2013 — ... : «Привет, альфа!».

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите год: ")
	year, _ := reader.ReadString('\n')
	yearInt, err := strconv.Atoi(year[:len(year)-1])

	if err != nil {
		fmt.Println(err)
		fmt.Println("Вы задали не год")
	}

	fmt.Println("Вы ввели год -", yearInt)

	switch {
	case yearInt <= 1964:
		fmt.Println("«Привет, бумер!»")
	case yearInt <= 1980:
		fmt.Println("«Привет, представитель X!»")
	case yearInt <= 1996:
		fmt.Println("«Привет, миллениал!»")
	case yearInt <= 2012:
		fmt.Println("«Привет, зумер!»")
	case yearInt >= 2013:
		fmt.Println("«Привет, альфа!»")
	default:
		fmt.Println("Не могу понять год :(")
	}
}
