package generators

import (
	"testing"

	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/mechanics"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/moons"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/planets"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/stars"
)

func TestApplyPlanetTidalLocking(t *testing.T) {
	testMoon := moons.Moon{
		MType: moons.MB,
		MOrbit: mechanics.Radius{
			Radius: 6,
			Unit:   mechanics.TacticalHex,
		},
		IsBig:        true,
		IsTideLocked: true,
		ResExpIndex:  mechanics.Normal,
	}
	testPlanetB := planets.Planet{
		Mass:  planets.MassTwo,
		Moons: []moons.Moon{testMoon},
		Orbit: mechanics.Orbit{
			Bearing: mechanics.Radian(2),
			Radius: mechanics.Radius{
				Radius: 1,
				Unit:   mechanics.LightMinute,
			},
		},
		PType:         planets.TypeB,
		HIndex:        0,
		Resexindex:    mechanics.Normal,
		TidallyLocked: true,
	}
	result := planets.Planet{
		Mass:  planets.MassTwo,
		Moons: []moons.Moon{testMoon},
		Orbit: mechanics.Orbit{
			Bearing: mechanics.Radian(2),
			Radius: mechanics.Radius{
				Radius: 1,
				Unit:   mechanics.LightMinute,
			},
		},
		PType:         planets.TypeT,
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

func TestGenerateMoon(t *testing.T) {
	testData := moons.Moon{
		MType: moons.MB,
		MOrbit: mechanics.Radius{
			Radius: 6,
			Unit:   mechanics.TacticalHex,
		},
		IsBig:        true,
		IsTideLocked: true,
		ResExpIndex:  mechanics.Normal,
	}
	moon := generateMoon(true, true, planets.TypeT, 7, 9, 3)
	if moon != testData {
		t.Error("Expected", testData, "received", moon)
	}
}
