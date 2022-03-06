package planets

import (
	"fmt"
	"testing"

	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/mechanics"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/stars"
)

func BenchmarkGetFormationZone(b *testing.B) {
	GetFormationZone(40, stars.OrangeStar)
}

func BenchmarkGetMass(b *testing.B) {
	GetMass(28)
}

func BenchmarkPlanetType(b *testing.B) {
	GetPlanetType(Lwz, MassTwo)
}

func ExampleGetFormationZone() {
	fmt.Println(GetFormationZone(10, stars.OrangeStar))
	//Output:
	//Gas false
}

func ExampleGetMass() {
	fmt.Println(GetMass(28))
	//Output:
	//Mass 2
}

func ExampleGetPlanetType() {
	fmt.Println(GetPlanetType(Lwz, MassTwo))
	//Output:
	//Type T
}

func TestGetFormationZone(t *testing.T) {
	for _, starData := range testData {
		for i, orbit := range starData.orbits {
			x, y := GetFormationZone(orbit, starData.startype)
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
		mass := GetMass(roll)
		if mass != results[i] {
			t.Error("Did not get the expected mass, got", mass, "expected", results[i])
		}
	}
}

func TestGetPlanetType(t *testing.T) {
	massData := []PlanetMass{Na, MassOne, MassTwo, MassThree}
	zoneData := []PlanetFormationZone{Hot, Lwz, Cold, Gas, Ice}
	results := [][]PlanetType{
		{AstB, TypeH, TypeV, TypeV},
		{AstB, TypeB, TypeT, TypeSt},
		{AstB, TypeB, TypeB, TypeB},
		{AstB, TypeB, TypeG, TypeG},
		{AstF, TypeF, TypeI, TypeI},
	}

	for i, zone := range zoneData {
		for j, mass := range massData {
			x := GetPlanetType(zone, mass)
			if x != results[i][j] {
				t.Error("Expected", results[i][j], "planet type but got", x)
			}
		}

	}
}
func TestDetermineRei(t *testing.T) {
	results := []mechanics.Rei{
		mechanics.VeryPoor,
		mechanics.Poor,
		mechanics.Poor,
		mechanics.Normal,
		mechanics.Normal,
		mechanics.Rich,
		mechanics.VeryRich}
	for i, num := range testRandomNumbers {
		rei := determineRei(num)
		if rei != results[i] {
			t.Error("Expected", results[i], "received", rei)
		}
	}
}

func TestGetReiNumber(t *testing.T) {
	num := 1
	results := []mechanics.Rei{
		mechanics.Normal,
		mechanics.Normal,
		mechanics.Poor,
		mechanics.Normal,
		mechanics.Normal,
		mechanics.VeryPoor,
		mechanics.VeryPoor,
		mechanics.VeryPoor}
	for i, planetType := range testPlanetTypes {
		n := GetReiNumber(planetType, num)
		if n != results[i] {
			t.Error("Expected", results[i], "received", n)
		}
	}
}
