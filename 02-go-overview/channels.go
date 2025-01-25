package main

import (
	"log"
	"math/rand"
)

//func PrintText(s string) {
//	log.Println(s)
//}

func RandomNumber(n int) int {
	value := rand.Intn(n) // he says you should seed with time.Now().UnixNano(),
				// but this is actually giving me random numbers wiothut using 
				// time.Now().UnixNano()
	return value
}

const numPool = 1000
func CalculateValue(intChan chan int) {
	randomNumber := RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	// PrintText("Hi")
	intChan := make(chan int)
	defer close(intChan)
	
	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
