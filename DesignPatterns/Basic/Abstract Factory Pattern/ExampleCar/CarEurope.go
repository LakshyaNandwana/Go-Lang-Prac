package main

type CarEurope struct{}

func (c *CarEurope) CreateCar() iCar {
	return &EUCar{
		Car: Car{
			carBrand: "Volkswagen",
			carType:  "Hatchback",
		},
	}
}

func (c *CarEurope) CreateSpec() iCarSpec {
	return &EUCarSpec{
		CarSpec: CarSpec{
			fuelCap:        "60L",
			safetyFeatures: "Europe Standard Safety Rating-5",
			milage:         18,
		},
	}
}
