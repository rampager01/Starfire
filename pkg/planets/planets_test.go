package planets

import (
	"fmt"
	"testing"

	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/mechanics"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/moons"
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
			x := getPlanetType(zone, mass)
			if x != results[i][j] {
				t.Error("Expected", results[i][j], "planet type but got", x)
			}
		}

	}
}

func TestGetNumberOfMoons(t *testing.T) {
	testData := []int{1, 15, 36, 67, 92, 115, 133}
	results := []int{1, 1, 1, 2, 3, 4, 5}
	for i, num := range testData {
		numMoons := getNumberOfMoons(num)
		if numMoons != results[i] {
			t.Error("Expeted", results[i], "received", numMoons)
		}
	}
}

func TestGetMoonType(t *testing.T) {
	results := []moons.MoonType{moons.MB, moons.MF, moons.MB, moons.MH, moons.MF, moons.MB, moons.MB, moons.MH}
	for i, planetType := range testPlanetTypes {
		moonType := getMoonType(planetType)
		if moonType != results[i] {
			t.Error("Expected", results[i], "recieved", moonType)
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
		n := getReiNumber(planetType, num)
		if n != results[i] {
			t.Error("Expected", results[i], "received", n)
		}
	}
}

func TestGenerateMoon(t *testing.T) {
	moon := generateMoon(true, true, TypeT, 7, 9, 3)
	if moon != testMoon {
		t.Error("Expected", testMoon, "received", moon)
	}
}

func TestApplyPlanetTidalLocking(t *testing.T) {
	result := Planet{
		Mass:  MassTwo,
		Moons: []moons.Moon{testMoon},
		Orbit: mechanics.Radius{
			Radius: 1,
			Unit:   mechanics.LightMinute,
		},
		PType:         TypeT,
		HIndex:        0,
		Resexindex:    mechanics.Normal,
		TidallyLocked: true,
	}
	moonResult := moons.Moon{
		MType: moons.MT,
		MOrbit: mechanics.Radius{
			Radius: 6,
			Unit:   mechanics.TacticalHex,
		},
		IsBig:        true,
		IsTideLocked: true,
		ResExpIndex:  mechanics.Normal,
	}
	ApplyPlanetTidalLocking(&testPlanetB, &testMoon, stars.RedDwarf)
	if testPlanetB.PType != result.PType && testMoon.MType != moonResult.MType {
		t.Error("Expected", result, moonResult, "received", testPlanetB, testMoon)
	}
}
