package main

import (
	"fmt"
)

// package level variable
// var myName string
// is allowed even if not used
// but he gets rid of it

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
	somethingMore, somethingYetMoreStill := saySomething2()

	fmt.Println("The function returned", whatWasSaid)
	fmt.Println("and finally, after", whatWasSaid, "was said, it was also said:", somethingMore, somethingYetMoreStill)
}

func saySomething() string {
	return "something" // first example
}

// possible to return more than one thing

func saySomething2() (string, string) {
	return "something", "and something else"
}
