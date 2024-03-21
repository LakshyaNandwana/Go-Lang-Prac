package creatures


type Werewolf struct{
	TimeofDay string
}


func (w Werewolf) Kind() string{
	if w.TimeofDay == "day"{
		return "human"
	}
	return "werewolf"
}

func (w Werewolf) Fly() bool{
	return false
}


func (w Werewolf) Sound() string{
	if w.TimeofDay == "day"{
		return "hello"
	}
	return "howl"
}


type Vampire struct{
	Age int
}

func (z Vampire) Kind() string{
	if z.Age > 100{
		return "Vampire"
	}
	return "human"
}

func (z Vampire) Fly() bool{
	return true
}

func (z Vampire) Sound() string{
	return "I want to drink your blood"
}