package main

import (
	"log"
)

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	
	changeUsingPointer(&myString) // using a reference

	log.Println("myString is now set to", myString)

	log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
