package main

import (
	"fmt"

	"halloween/creatures"
	"halloween/laboratory"

)


func makeVampire(age int) laboratory.Creatures{
	return creatures.Vampire{Age: age}
}


func main(){

	raman := creatures.Werewolf{TimeofDay: "night"}

	fmt.Printf("Raman is a %s\n", raman.Kind())
	fmt.Println(laboratory.Greet(raman))

	fmt.Println(laboratory.FlyAway(raman))

	raman.TimeofDay = "day"

	fmt.Printf("Raman is %s\n ", raman.Kind())

	fmt.Println(laboratory.Greet(raman))
	fmt.Println(laboratory.FlyAway(raman))


	Rohit := makeVampire(132)

	fmt.Printf("Rohit is a %s \n", Rohit.Kind())

	fmt.Println(laboratory.Greet(Rohit))
	fmt.Println(laboratory.FlyAway(Rohit))


}