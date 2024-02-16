package main

import "fmt"

type iCarFactory interface{
	CreateCar() iCar
	CreateSpec() iCarSpec
}


func getCarFactory(carType string) (iCarFactory, error){
	if carType == "SUV"{
		return &CarNorthAmerica{}, nil
	}
	if carType == "Hatchback"{
		return &CarEurope{}, nil
	}
	return nil, fmt.Errorf("Invalid input of car")
}