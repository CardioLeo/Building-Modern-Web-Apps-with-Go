package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World.")

	var whatToSay string

	whatToSay = "Goodbye, cruel world"

	fmt.Println(whatToSay)

	whatToSay = "Sorry, I meant: 'Goodbye, relatively cruel world'"

	fmt.Println(whatToSay)
}
