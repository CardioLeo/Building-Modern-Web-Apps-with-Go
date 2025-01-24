package main

import (
	"log"
)

// "something you almost never see in Go - a semi-colon"

func main() {
	for i := 0; i <= 10; i++ { // passing these "to the for keyword"
		log.Println(i)
	}

	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}
}
