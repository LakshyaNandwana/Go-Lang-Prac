package main

import "fmt"

func main() {
    // Declare a variable
    var num int = 42

    // Declare a pointer variable and assign the address of 'num' to it
    var ptr *int = &num

    // Print the value and address of 'num'
    fmt.Println("Value of num:", num)
    fmt.Println("Address of num:", &num)

    // Print the value and address stored in the pointer 'ptr'
    fmt.Println("Value stored in ptr:", *ptr)
    fmt.Println("Address stored in ptr:", ptr)

    // Modify the value indirectly using the pointer
    *ptr = 100
    fmt.Println("Modified value of num:", num)
}
