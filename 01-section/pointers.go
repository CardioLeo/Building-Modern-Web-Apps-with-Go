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
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
	log.Println("s is now set to", s)
}
