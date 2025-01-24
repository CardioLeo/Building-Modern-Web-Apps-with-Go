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

	for _, animal := range animals {
		log.Println(animal)
	}

	animals2 := make(map[string]string)
	animals2["cat"] = "Fluffy"
	animals2["dog"] = "Fido"

	for animalType, animal := range animals2 {
		log.Println(animalType, animal)
	}

	var firstLine = "Once upon a midnight dreary"
	for i, l := range firstLine {
		log.Println(i, ":", l)
		log.Printf("%c:%c", i, l)
	}

}
