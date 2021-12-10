package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	engBot := englishBot{}
	spBot := spanishBot{}

	printGreeting(engBot)
	printGreeting(spBot)
}

func printGreeting(myBot bot) {
	fmt.Println(myBot.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
