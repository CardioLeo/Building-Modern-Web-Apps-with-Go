package main

import (
	"log"
	// "github.com/cardioleo/myniceprogram/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	log.Println(myVar.TypeName)
}

// go modules is how we use go packages now; work in IDE

// go mod init name-of-module
// go mod init github.com/tsawler/myniceprogram

// "sometimes I have to fight with goland to get it to recognize it." ugh - this was my
// I had no luck
// have to enable it in GoLand; maybe not in VSCode -- what about VIM???

// yeah, this doesn't work; I figure you have to make it work in an IDE. Kind of stupid.
