package main

import(
	"log"
	"errors"
)

func divide(x, y float32) (float32, error) {
	// takes two parameters and returns two
	var result float32
	result = x / y
	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}
	return result, nil
}

func main() {
	result, err := divide(100.0, 10.0)
	if err != nil {
                log.Println(err)
                return
        }
	log.Println("result of division is", result)
}
