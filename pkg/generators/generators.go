package generators

import (
	"math/rand"
	"time"

	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/mechanics"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/moons"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/planets"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/stars"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/warppoints"
)

func GenerateWarpPoint(bearing, radius, wpType, visiblity int) warppoints.WarpPoint {
	orbit := mechanics.Orbit{
		Bearing: mechanics.Radian(bearing),
		Radius: mechanics.Radius{
			Radius: radius,
			Unit:   mechanics.SystemHex,
		},
	}
	return warppoints.WarpPoint{
		Orbit:          orbit,
		WPType:         warppoints.GetWarpPointType(wpType),
		Visibility:     warppoints.GetWarpPointVisibility(visiblity),
		WPSystemLinkID: 0,
	}
}

func GeneratePlanet(orbit int, starType stars.StarType, s rand.Source) planets.Planet {
	r := rand.New(s)
	ranNum := r.Intn(11) + 1
	orbitalRadius := mechanics.Orbit{
		Bearing: mechanics.Radian(ranNum),
		Radius: mechanics.Radius{
			Radius: orbit,
			Unit:   mechanics.LightMinute},
	}
	r = rand.New(s)
	ranNum = r.Intn(99) + 1
	mass := planets.GetMass(ranNum)

	zone, isLocked := planets.GetFormationZone(orbitalRadius.Radius.Radius, starType)

	planetType := planets.GetPlanetType(zone, mass)

	ranNum = r.Intn(99) + 1
	m := generateMoons(planetType, mass, ranNum)

	index := 0
	if planetType == planets.TypeT || planetType == planets.TypeSt {
		index = r.Intn(9) + 1
	}

	ranNum = r.Intn(9) + 1
	resExIndex := planets.GetReiNumber(planetType, ranNum)

	return planets.Planet{
		Mass:          mass,
		Moons:         m,
		Orbit:         orbitalRadius,
		PType:         planetType,
		HIndex:        index,
		Resexindex:    resExIndex,
		TidallyLocked: isLocked,
	}
}

func generateMoons(planetType planets.PlanetType, mass planets.PlanetMass, num int) []moons.Moon {
	numMoons := 0
	switch planetType {
	case planets.TypeB, planets.TypeT, planets.TypeSt:
		{
			switch mass {
			case planets.MassOne:
				{
					numMoons = getNumberOfMoons(num - 50)
				}
			case planets.MassTwo:
				{
					numMoons = getNumberOfMoons(num - 10)
				}
			case planets.MassThree:
				{
					numMoons = getNumberOfMoons(num)
				}
			}
		}
	case planets.TypeH:
		{
			switch mass {
			case planets.MassOne:
				{
					numMoons = getNumberOfMoons(num - 65)
				}
			case planets.MassTwo:
				{
					numMoons = getNumberOfMoons(num - 25)
				}
			case planets.MassThree:
				{
					numMoons = getNumberOfMoons(num - 15)
				}
			}
		}
	case planets.TypeF:
		{
			numMoons = getNumberOfMoons(num - 15)
		}
	case planets.TypeV:
		{
			numMoons = getNumberOfMoons(num - 35)
		}
	case planets.TypeI:
		{
			numMoons = getNumberOfMoons(num + 35)
		}
	case planets.TypeG:
		{
			numMoons = getNumberOfMoons(num + 50)
		}
	}

	planetMoons := make([]moons.Moon, numMoons)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	a := 0
	priorDistance := 0
	for i := 0; i < numMoons; i++ {
		isBig := false
		isLocked := false
		if i == 0 {
			isBig, isLocked = moons.CheckTideLocked(numMoons, 0)
		}
		a = mechanics.FractionRoundUp(r.Intn(9)+1, 2)
		reiNum := r.Intn(9) + 1
		planetMoons = append(planetMoons, generateMoon(isBig, isLocked, planetType, reiNum, a, priorDistance))
		priorDistance = planetMoons[i].MOrbit.Radius
	}
	return planetMoons
}

func getNumberOfMoons(rand int) int {
	if rand >= 1 && rand < 56 {
		return 1
	}
	if rand >= 56 && rand < 86 {
		return 2
	}
	if rand >= 86 && rand < 106 {
		return 3
	}
	if rand >= 106 && rand < 127 {
		return 4
	}
	if rand >= 127 {
		return 5
	}
	return 0
}

func generateMoon(isBig, isLocked bool, planetType planets.PlanetType, reiNum, num, priorDistance int) moons.Moon {
	return moons.Moon{
		MType: getMoonType(planetType),
		MOrbit: mechanics.Radius{
			Radius: getMoonDistance(priorDistance, num, planetType),
			Unit:   mechanics.TacticalHex},
		IsBig:        isBig,
		IsTideLocked: isLocked,
		ResExpIndex:  planets.GetReiNumber(planetType, reiNum),
	}
}

func getMoonType(planetType planets.PlanetType) moons.MoonType {
	switch planetType {
	case planets.TypeB, planets.TypeT, planets.TypeSt, planets.TypeG:
		{
			return moons.MB
		}
	case planets.TypeH, planets.TypeV:
		{
			return moons.MH
		}
	case planets.TypeI, planets.TypeF:
		{
			return moons.MF
		}
	}
	return moons.MB
}

func getMoonDistance(priorDistance, randNum int, planetType planets.PlanetType) int {
	switch planetType {
	case planets.TypeB, planets.TypeF, planets.TypeH, planets.TypeSt, planets.TypeT, planets.TypeV:
		{
			priorDistance = priorDistance + mechanics.FractionRoundUp(randNum, 3)
		}
	case planets.TypeG, planets.TypeI:
		{
			if priorDistance == 0 {
				priorDistance = mechanics.FractionRoundUp(randNum, 2)
			} else {
				priorDistance = priorDistance + randNum
			}
		}
	}
	return priorDistance
}

func ApplyPlanetTidalLocking(planet *planets.Planet, moon *moons.Moon, starType stars.StarType) {
	switch starType {
	case stars.WhiteStar, stars.YellowWhiteStar, stars.YellowStar, stars.OrangeStar, stars.RedStar:
		{
			if planet.TidallyLocked {
				switch planet.Mass {
				case planets.MassTwo, planets.MassThree:
					{
						planet.PType = planets.TypeV
					}
				case planets.MassOne:
					{
						planet.PType = planets.TypeH
						moon.MType = moons.MH
					}
				}
			} else {
				if moon.IsTideLocked && planet.PType == planets.TypeG {
					moon.MType = moons.MH
				}
			}
		}
	case stars.RedDwarf:
		{
			if planet.TidallyLocked {
				switch planet.Mass {
				case planets.MassOne:
					{
						planet.PType = planets.TypeH
						moon.MType = moons.MT
					}
				case planets.MassTwo:
					{
						if moon.IsBig {
							planet.PType = planets.TypeT
							moon.MType = moons.MT
						} else {
							planet.PType = planets.TypeH
							moon.MType = moons.MB
						}
					}
				case planets.MassThree:
					{
						if moon.IsBig {
							planet.PType = planets.TypeSt
							moon.MType = moons.MT
						} else {
							planet.PType = planets.TypeH
							moon.MType = moons.MB
						}
					}
				}
			} else {
				moon.MType = moons.MH
			}
		}
	}
}
