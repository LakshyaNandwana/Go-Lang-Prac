package main

import "fmt"

type iSportFactory interface{
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsFactory(brand string) (iSportFactory, error){
	if brand == "adidas"{
		return &adidas{}, nil
	}
	if brand =="nike"{
		return &nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type  passed")
}
