package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World.")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"

	fmt.Println(whatToSay)

	whatToSay = "Sorry, I meant: 'Goodbye, relatively cruel world'"

	fmt.Println(whatToSay)

	// i = "cat" // doesn't work, he says, and I'm following along

	i = 7

	fmt.Println("i is set to", i)
}
