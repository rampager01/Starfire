package planets

import (
	"fmt"
	"testing"

	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/stars"
)

func BenchmarkGetFormationZone(b *testing.B) {
	getFormationZone(40, stars.OrangeStar)
}

func BenchmarkGetMass(b *testing.B) {
	getMass(28)
}

func BenchmarkPlanetType(b *testing.B) {
	getPlanetType(Lwz, MassTwo)
}

func ExamplegetFormationZone() {
	fmt.Println(getFormationZone(10, stars.OrangeStar))
	//Output:
	//Gas false
}

func ExamplegetMass() {
	fmt.Println(getMass(28))
	//Output:
	//MassTwo
}

func ExamplegetPlanetType() {
	fmt.Println(getPlanetType(Lwz, MassTwo))
	//Output:
	//TypeT
}

func TestGetFormationZone(t *testing.T) {
	for _, starData := range testData {
		for i, orbit := range starData.orbits {
			x, y := getFormationZone(orbit, starData.startype)
			if x != starData.results[i].zone {
				t.Error("Planetary formation zones does not match, Got", x, "expected ", starData.results[i].zone)
			}
			if y != starData.results[i].tideLocked {
				t.Error("Tide Locked zone does not match, Got", y, "expected ", starData.results[i].tideLocked)
			}
		}
	}
}

func TestGetMass(t *testing.T) {
	testData := []int{1, 5, 18, 75, 99}
	results := []PlanetMass{None, Na, MassOne, MassTwo, MassThree}
	for i, roll := range testData {
		mass := getMass(roll)
		if mass != results[i] {
			t.Error("Did not get the expected mass, got", mass, "expected", results[i])
		}
	}
}

func TestGetPlanetType(t *testing.T) {
	massData := []PlanetMass{Na, MassOne, MassTwo, MassThree}
	zoneData := Lwz //[]PlanetFormationZone{Hot, Lwz, Cold, Gas, Ice}
	results := []PlanetType{AstB, TypeB, TypeT, TypeSt}
	for i, mass := range massData {
		x := getPlanetType(zoneData, mass)
		if x != results[i] {
			t.Error("Expected", results[i], "planet type but got", x)
		}
	}
}
