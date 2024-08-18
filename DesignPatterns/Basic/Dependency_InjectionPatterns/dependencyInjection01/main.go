package main

import "fmt"

type SafetyPlacer interface {
	placeSafeties()
}
type RockClimber struct {
	kind         int
	rocksClimbed int
	sp           SafetyPlacer
}

func NewRockClimber(sp SafetyPlacer) *RockClimber {
	return &RockClimber{
		sp: sp,
	}
}

type IceSafetyPlaces struct {
	//Database
	//data
	//api
}

func (sp IceSafetyPlaces) placeSafeties() {
	fmt.Println("Placing safeties for iccy rocks")
}

type NOPSafetyPlacer struct{}

func (sp NOPSafetyPlacer) placeSafeties() {
	fmt.Println("Placing no safeties....")
}

func (rc *RockClimber) climbRock() {
	rc.rocksClimbed++
	if rc.rocksClimbed == 10 {
		rc.sp.placeSafeties()
	}
}

func main() {

	rc := NewRockClimber(IceSafetyPlaces{})
	for i := 0; i < 11; i++ {
		rc.climbRock()
	}

	rc = NewRockClimber(NOPSafetyPlacer{})
	for i := 0; i < 11; i++ {
		rc.climbRock()
	}
}
