package main

type iCar interface{
	setBrand(carBrand string)
	setCarType(carType string)
	getBrand() string
	getCarType() string
}

type Car struct{
	carBrand string
	carType string
}

func(c *Car) setBrand(carBrand string){
	c.carBrand = carBrand
}

func(c *Car) setCarType(carType string){
	c.carType = carType
}

func(c *Car) getBrand() string{
	return c.carBrand
}

func(c *Car) getCarType() string{
	return c.carType
}