package main

import "fmt"

func main()  {
	type Name string
	type Fruit string

	var fruit Fruit
	var name Name

	fruit = "Apple"
	name = Name(fruit) // Явно приводим тип

	fmt.Println("name, fruit", fruit, name)
}

