package main

import (
	"fmt"
)

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberOfTeeth int
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func main() {
	dog := Dog {
		Name: "Samson",
		Breed: "German Shepherd",
	}

	PrintInfo(dog)
}
