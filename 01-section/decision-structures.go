package main

import (
	"log"
)

func main() {
	// how to make decisions in a program, in Go
	// very like other languages

	var isTrue bool
	isTrue = true
	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}
}
