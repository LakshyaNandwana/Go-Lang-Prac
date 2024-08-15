package main

import "fmt"

type Key byte

const (
	Silver Key = iota + 1
	Gold
	Diamond
)

func (k Key) String() string {
	switch k {
	case Silver:
		return "silver"

	case Gold:
		return "gold"

	case Diamond:
		return "diamond"

	}
	return fmt.Sprintf("<Key %d>", k)
}

type Treasure struct {
	Name string
	Keys []Key
}
