package main

import (
	"encoding/json"
	"log"
)


type Person Struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog bool `json:"has_dog"`
}

func main() {

	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshall([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalled json", err)
	}
	
	log.Printf("unmarshalled: %v", unmarshalled)
}
