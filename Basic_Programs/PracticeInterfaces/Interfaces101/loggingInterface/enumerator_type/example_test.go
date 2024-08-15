package main

import "fmt"

func ExampleKey() {
	fmt.Println(Silver)
	fmt.Println(Key(9))

	fmt.Printf(" v: %v\n", Gold)
	fmt.Printf("+v: %+v\n", Gold)
	fmt.Printf("#v: %#v\n", Gold)

	t1 := Treasure{
		Name: "The Amber Room",
		Keys: []Key{Gold},
	}

	fmt.Printf("t1: %#v\n", t1)

	// Output:
	// silver
	// <Key 9>
	//  v: gold
	// +v: gold
	// #v: 0x2
	// t1: main.Treasure{Name:"The Amber Room", Keys:[]main.Key{0x2}}

}
