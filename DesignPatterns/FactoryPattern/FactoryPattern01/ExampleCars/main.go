package main

import "fmt"

type Car interface{
	Drive() string
}

type Sedan struct{}
type SUV struct{}

func(s *Sedan) Drive() string{
	return "Driving Sedan car"
}

func(s *SUV) Drive() string{
	return "Driving SUV car"
}

type CarFactory interface{
	CreateCar() Car
}

type SedanFacotry struct{}
type SUVFactory struct{}

func(sf *SedanFacotry) CreateCar() Car{
	return &Sedan{}
}

func (sf *SUVFactory) CreateCar() Car{
	return &SUV{}
}

func main(){
	sedanFactory := &SedanFacotry{}
	sedan := sedanFactory.CreateCar()
	fmt.Println(sedan.Drive())

	suvFactory := &SUVFactory{}
	suv := suvFactory.CreateCar()
	fmt.Println(suv.Drive())
}