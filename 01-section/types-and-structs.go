package main

import (
	"log"
)

var s = "seven"

var firstName string
var lastname string
var phoneNumber string
var age int
var birthDate time.Time // take advantage of strong typing

func main() {
	var s2 = "six"

	s := "eight"

	log.Println(s) // pulls on local var, not package level var
	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("xxx") // pulls on package level var, not local var
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething func is", s)
	return s3, "World"
}
