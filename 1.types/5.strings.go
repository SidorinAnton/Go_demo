package main

import "fmt"


func main()  {
	var text string
	text = "Hello"

	fmt.Println("text -", text)
	fmt.Println("text[0] -", text[0])
	fmt.Println("text + text -", text + text)
	fmt.Println("len(text) -", len(text))

	var text2 = `
	Hello
	World
	`
	fmt.Println("`text` -", text2)




}

