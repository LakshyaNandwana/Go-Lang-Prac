package main

import "fmt"

// Define a base struct with a single field 'num'
type base struct {
    num int
}

// Method for the 'base' struct that describes its content
func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

// Define a container struct that embeds the 'base' struct and has an additional field 'str'
type container struct {
    base
    str string
}

func main() {
    // Create an instance of the 'container' struct with values for both 'base' and 'str'
    co := container{
        base: base{
            num: 1,
        },
        str: "some name",
    }

    // Print the values of 'num' and 'str' in the 'container'
    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

    // Access 'num' directly from the embedded 'base' struct
    fmt.Println("also num:", co.base.num)

    // Invoke the 'describe' method on the 'container', which was inherited from 'base'
    fmt.Println("describe:", co.describe())

    // Create an interface 'describer' that requires a 'describe' method
    type describer interface {
        describe() string
    }

    // Assign the 'container' to the 'describer' interface
    var d describer = co

    // Call the 'describe' method through the interface
    fmt.Println("describer:", d.describe())
}
