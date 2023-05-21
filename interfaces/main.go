package main

import "fmt"

type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// very custom logic for generating an spanish greeting
	return "Hola!"
}


