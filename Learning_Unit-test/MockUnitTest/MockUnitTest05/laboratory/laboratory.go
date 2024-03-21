package laboratory

import "fmt"


type Creatures interface{
	Kind() string
	Fly() bool
	Sound() string
}


func Greet(c Creatures) string{
	return c.Sound()
}


func FlyAway(canfly Creatures) string{
	if canfly.Fly() == true{
		return fmt.Sprintf("%s's can fly", canfly.Kind())
	}
	return fmt.Sprintf("%s's cannot fly", canfly.Kind())
}