package main

//Interfaces Assignement
import "fmt"

type engilishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func d() {

	eb := engilishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (eb engilishBot) getGreeting() string {
	//Very custoom logic
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
