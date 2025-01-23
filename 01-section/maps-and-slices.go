package main

import (
	"log"
)

func main() {
	var myString string
	var myInt int

	myString = "Hi"
	myInt = 11

	mySecondString := "another string"

	log.Println(myString, mySecondString, myInt)

	myMap := make(map[string]string)
	// var myOtherMap map[string]string // can't do anything with this; or only rarely
					// he has yet to encounter this

	myMap["dog"] = "Samson"

	log.Println(myMap["dog"])
}
