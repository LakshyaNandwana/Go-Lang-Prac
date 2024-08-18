//go:build wireinject
// +build wireinject

// wire.go
package main

import "github.com/google/wire"

func InitializeEvent(phase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMesssage)

	return Event{}, nil
}

//Read to understand
/*Rather than go through the trouble of initializing each component in turn and passing it into the next one, we instead have a single call
to wire.Build passing in the initializers we want to use.

We add a zero value for Event as a return value to satisfy the compiler.
* Note that even if we add values to Event, Wire will ignore them.


*/
