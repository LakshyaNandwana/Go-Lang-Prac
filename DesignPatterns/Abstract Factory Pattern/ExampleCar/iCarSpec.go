package main

type iCarSpec interface{
	setFuelCapacity(fuelCap string)
	setSafetyFeatures(safetyFeatures string)
	setMilage(milage int)
	getFuelCapacity() string
	getSafetyFeatures() string
	getMilage() int
}

type CarSpec struct{
	fuelCap string
	safetyFeatures string
	milage int
}

func(c *CarSpec) setFuelCapacity(fuelCap string){
	c.fuelCap = fuelCap
}

func(c *CarSpec) setSafetyFeatures(safetyFeatures string){
	c.safetyFeatures = safetyFeatures
}

func(c *CarSpec) setMilage(milage int){
	c.milage = milage
}

func(c *CarSpec) getFuelCapacity() string{
	return c.fuelCap
}

func(c *CarSpec) getSafetyFeatures() string{
	return c.safetyFeatures
}

func(c *CarSpec) getMilage() int{
	return c.milage
}
