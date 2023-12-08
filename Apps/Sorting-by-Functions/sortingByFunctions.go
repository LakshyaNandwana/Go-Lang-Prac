package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	//Sorting People based on Age
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)

	// nameCmp := func(a,b Person) int{
	// 	return cmp.Compare(a.name, b.name)
	// }
	//Sorting People based on Name
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.name, b.name)
		})
	fmt.Println(people)
}
