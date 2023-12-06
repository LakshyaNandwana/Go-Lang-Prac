package main

import (
	"fmt"
	"errors"
)

func plus(a int, b int) int{
	return a+b
}

func plusPlus(a,b,c int)int{
	return a+b+c
}

//Function that return multiple values
func vals() (int, int){
	return 3,7
}

//Function that returns result and errors
func divide(a, b float64) (float64, error){
	if b==0{
		//Returns error if divisor is zero
		return 0, errors.New("Cannot divide by zero")
	}

	return a/b, nil
}



func main(){
	res := plus(1,2)
	fmt.Println("1+2=", res)

	res = plusPlus(1,2,3)
	fmt.Println("1+2+3=", res)


	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
	
	d, e := vals()
	fmt.Println(d,e)

	result , err := divide(10,2)
	if err != nil{
		fmt.Println("Error:", err)
	} else{
		fmt.Println("Result:", result)
	}

	result , err = divide(10,0)
	if err != nil{
		fmt.Println("Error:", err)
	} else{
		fmt.Println("Result:", result)
	}
}