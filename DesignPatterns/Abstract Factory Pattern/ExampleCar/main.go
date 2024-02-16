package main

import "fmt"

func main(){
	northAmericaCarFactory, _ := getCarFactory("SUV")
	europeCarFactory, _ := getCarFactory("Hatchback")
	naCar := northAmericaCarFactory.CreateCar()
	naCarSpec := northAmericaCarFactory.CreateSpec()
	euCar := europeCarFactory.CreateCar()
	euCarSpec := europeCarFactory.CreateSpec()

	printCarDetails(naCar)
	printCarDetails(euCar)

	printCarSpecDetails(naCarSpec)
	printCarSpecDetails(euCarSpec)
}


func printCarDetails(c iCar){
	fmt.Printf("Brand: %s\n", c.getBrand())
	fmt.Printf("Type: %s\n", c.getCarType())
}


func printCarSpecDetails(c iCarSpec){
	fmt.Printf("Fuel Capacity: %s\n", c.getFuelCapacity())
	fmt.Printf("Safety Features: %s\n", c.getSafetyFeatures())
	fmt.Printf("Milage: %d\n", c.getMilage())
}