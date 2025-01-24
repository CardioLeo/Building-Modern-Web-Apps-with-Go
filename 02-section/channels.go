package main

import (
	"log"
	"math/rand"
)

func PrintText(s string) {
	log.Println(s)
}

func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}

const numPool = 10
func CalculateValue(intChan chan int) {
	randomNumber := RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	PrintText("Hi")
	intChan := make(chan int)
	defer close(intChan)
}
