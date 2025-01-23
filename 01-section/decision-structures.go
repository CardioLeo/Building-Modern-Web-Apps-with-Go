package main

import (
	"log"
)

func main() {
	// how to make decisions in a program, in Go
	// very like other languages

	var isTrue bool
	// isTrue = false
	isTrue = true
	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("Cat is", cat)
	} else {
		log.Println("Cat is not", cat)
	}

	myNum := 100
	isTrue2 := false

	if (myNum > 99 && isTrue) { // deliberately not isTrue2
		log.Println("myNum is greater than 99 and isTrue is set to", isTrue, "but isTrue2 is set to", isTrue2)
	}

	// he has a bunch more else conditionals - but I'm tired, it's late,
	// and I've done this before. I'm going over this aspect mostly just for practice

	// switch statements

	myVar := "cat"
	switch myVar {
		case "cat":
			log.Println("cat is set to cat")
		case "dog":
			log.Println("cat is set to dog")
		case "fish":
			log.Println("cat is set to fish")
		default:
			log.Println("cat is something else")
	}
}
