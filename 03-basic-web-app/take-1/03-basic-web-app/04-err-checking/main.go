package main

import (
	"fmt"
	"net/http"
	"errors"
)


const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	// "This is going to be a handler function;" has to handle two parameters
	fmt.Fprintf(w, "This is the home page")
}

// first thing that should appear before function is the function name,
// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
	// adds underscore to beginning of line because I don't care about the error
	// which this function will return
}

// addValues adds two integers and returns the sum
func addValues(x, y int) (int) {
	sum := x + y
	return sum
}

func divideValues(x, y float32) (float32, error){
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x, y float32 = 100.0, 0.0
	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by zero")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}

// main is the main application function
func main() {
	// fmt.Println("Hello, world")
	/*
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// first parameter is the url that I want this to listen to
		n, err := fmt.Fprintf(w, "Hello, world!")
		// fmt.Fprintf(w, "Hello, world!")
		// fmt.Println("Bytes written:" + n)
		// "n is an int, and you can't print an int like that" - dude is being Dr Seuss
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})
	*/
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	// so now, you can go to this by putting "http://localhost:8080/about" into the
	// search bar / address bar

	// "how easy it is to start a webserver that listens in go"
	// _ = http.ListenAndServe(":8080", nil) // may be preferable

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

	// starting with the underscore like that says, "if there's an error, I don't care"
	// important comment because the underscore is so confusing if you don't know!
	// "this is the basis for any web application in go" very cool

	// if you go to "http://localhost:8080/" in a web browser while
	// you're running this command, you will see hello world. (Works.)
	// Also it does print out number of bytes twice every time that
	// you refresh. He comments on how he doesn't know why and isn't
	// worried about it right now. Mine does the same thing.
}
