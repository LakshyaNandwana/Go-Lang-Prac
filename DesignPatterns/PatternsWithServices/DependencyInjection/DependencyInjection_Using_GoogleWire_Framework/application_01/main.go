package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Message string

type Greeter struct {
	Grumpy  bool
	Message Message
}

type Event struct {
	Greeter Greeter
}

func NewMesssage(phase string) Message {
	return Message(phase)
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// Defining the above componentes for application.
// Initialize all the components without using Wire. Main looks like:-
/*
func main() {
	message := NewMesssage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
}

*/
//Using Wire to generate code
func main() {
	e, err := InitializeEvent("Hi there !")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}

	e.Start()
}
