package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName string
}

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
	myMap["other-dog"] = "Cassie"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap["dog"] = "fido"

	log.Println("now my dog's name is", myMap["dog"])

	myMap2 := make(map[string]int)
	myMap2["First"] = 1
	myMap2["Second"] = 2

	log.Println(myMap2["First"], myMap2["Second"])
	
	myMap3 := make(map[string]User)
	
	me := User {
		FirstName: "Trevor",
		LastName: "Sawler",
	}

	myMap3["me"] = me
	log.Println(myMap3["me"].FirstName)

	// var myNewVar float32

	// myNewVar = 11.1

	// myMap4 := make(map[string]interface{}) // not recommended -> map of interfaces, I guess
			// will store anything you want
			// but have to cast it (type cast) to what you need

	// onto slices
}
