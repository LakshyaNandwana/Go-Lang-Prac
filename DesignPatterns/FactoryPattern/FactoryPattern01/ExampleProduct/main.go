package main

import "fmt"

type Product interface{
	GetName() string
}

type Factory interface{
	CreateProduct() Product
}

type ConcreteProduct struct{
	name string
}

func(p *ConcreteProduct) GetName() string{
	return p.name
}

type ConcreteFactory struct{}

func (f *ConcreteFactory) CreateProduct() Product{
	return &ConcreteProduct{name: "Concrete Product"}
}

func main(){
	factory := &ConcreteFactory{}
	product := factory.CreateProduct()
	fmt.Println(product.GetName())
}