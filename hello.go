package main

import (
	"fmt"
	"greetings"
	"log"

	"rsc.io/quote"
)

func main() {
	//Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source fie and line nuber
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	//Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	// message, err := greetings.Hello("Gladys")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(message)
	fmt.Println(quote.Go())
}
