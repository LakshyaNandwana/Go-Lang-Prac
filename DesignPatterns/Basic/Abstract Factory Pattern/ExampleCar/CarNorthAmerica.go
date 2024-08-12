package main

type CarNorthAmerica struct{}

func (c *CarNorthAmerica) CreateCar() iCar {
	return &NACar{
		Car: Car{
			carBrand: "BMW",
			carType:  "Sedan",
		},
	}
}


func(c *CarNorthAmerica) CreateSpec() iCarSpec{
	return &NACarSpec{
		CarSpec : CarSpec{
			fuelCap : "45L",
			safetyFeatures: "Mandatory Airbags",
			milage: 10,	
		},
	}
}