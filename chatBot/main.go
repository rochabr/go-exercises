package main

import (
	"fmt"
	"os"
)

type englishBot struct{}

type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	// eb := englishBot{}
	// sb := spanishBot{}

	// printGreeting(eb)
	// printGreeting(sb)

	// t := triangle{
	// 	height: 10,
	// 	base:   5,
	// }

	// s := square{
	// 	sideLength: 10,
	// }

	// printArea(t)
	// printArea(s)

	fmt.Println(os.Args[1])
	readFile(os.Args[1])
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "hello"
}

func (spanishBot) getGreeting() string {
	return "hola"
}
