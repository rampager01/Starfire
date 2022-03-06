package modifiers

import (
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/mechanics"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/moons"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/planets"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/warppoints"
)

func WarpPointPlanetCollisionCheck(planet *planets.Planet, warpPoints []warppoints.WarpPoint) {
	for _, warpPoint := range warpPoints {
		if warpPoint.Orbit.Radius.Radius == planet.Orbit.Radius.Radius {
			switch planet.PType {
			case planets.TypeB, planets.TypeH, planets.TypeV, planets.TypeT, planets.TypeSt, planets.TypeG:
				{
					planet.PType = planets.AstB
					planet.Moons = []moons.Moon{}
					planet.Resexindex = mechanics.Rich
				}
			case planets.TypeF, planets.TypeI:
				{
					planet.PType = planets.AstF
					planet.Moons = []moons.Moon{}
					planet.Resexindex = mechanics.Normal
				}
			}
		}
	}
}

func GasGiantDisruptionCheck(randNum int, planetList []planets.Planet) {
	for i := len(planetList) - 1; i >= 0; i-- {

	}
}
