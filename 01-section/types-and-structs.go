package main

import (
	"log"
	"time"
)

var s = "seven"

var firstName string
var lastname string
var phoneNumber string
var age int
var birthDate time.Time // take advantage of strong typing

type User struct {
	// this struct contains all sorts of information
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {
	var s2 = "six"

	s := "eight"

	log.Println(s) // pulls on local var, not package level var
	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("xxx") // pulls on package level var, not local var

	// now that struct is created, code below

	user := User {
		FirstName: "Trevor",
		LastName: "Sawler",
		PhoneNumber: "blargidy-blarg",
		Age: 35,
		// BirthDate:
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething func is", s)
	return s3, "World"
}
