package main

import "fmt"


type dataParser interface{
	toString() string
}

type robot struct{
	name string
	kind string
	autonomy float32
}

func (r robot) toString() string{
	return fmt.Sprintf("Name: %s\n"+
	"Kind: %s\n"+
	"Autonomy: %2.2f",
	r.name,
	r.kind,
	r.autonomy)
}


func newRobot(name, kind string, autonomy float32) robot{
	return robot{
		name: name,
		kind: kind,
		autonomy: autonomy,
	}
}



func main(){

	teacherRobot := newRobot("Mr. teacher Robot","teacher", 90)
	fighterRobot := newRobot("Mr. fighter Robot","fighter",70)

	fmt.Println(teacherRobot.toString())
	fmt.Println(fighterRobot.toString())

}