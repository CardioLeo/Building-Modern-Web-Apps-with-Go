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

	test := saySomething()
	fmt.Println(test)
	whatWasSaid := saySomething()
	fmt.Println(whatWasSaid)

	fmt.Println("The function returned", whatWasSaid)
}

func saySomething() string {
	return "something" // first example
}
