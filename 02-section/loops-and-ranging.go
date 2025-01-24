package main

import (
	"log"
)

// "something you almost never see in Go - a semi-colon"

func main() {
	for i := 0; i <= 10; i++ { // passing these "to the for keyword"
		log.Println(i)
	}
}
