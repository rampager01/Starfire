package planets

import (
	"math/rand"
	"time"

	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/mechanics"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/moons"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/stars"
)

type PlanetMass string
type PlanetType string
type PlanetFormationZone string

const (
	None      PlanetMass = "No Planet"
	Na        PlanetMass = "NA"
	MassOne   PlanetMass = "Mass 1"
	MassTwo   PlanetMass = "Mass 2"
	MassThree PlanetMass = "Mass 3"
)

const (
	AstB   PlanetType = "AST-B"
	AstF   PlanetType = "AST-F"
	No     PlanetType = "No Planet"
	TypeB  PlanetType = "Type B"
	TypeF  PlanetType = "Type F"
	TypeG  PlanetType = "Type G"
	TypeH  PlanetType = "Type H"
	TypeI  PlanetType = "Type I"
	TypeT  PlanetType = "Type T"
	TypeSt PlanetType = "Type ST"
	TypeV  PlanetType = "Type V"
)

const (
	Hot        PlanetFormationZone = "Hot"
	Lwz        PlanetFormationZone = "LWZ"
	Cold       PlanetFormationZone = "Cold"
	Gas        PlanetFormationZone = "Gas"
	Ice        PlanetFormationZone = "Ice"
	TideLocked PlanetFormationZone = "Tidally Locked"
)

type Planet struct {
	Mass          PlanetMass
	Moons         []moons.Moon
	Orbit         mechanics.Radius
	PlanetType    PlanetType
	HIndex        int
	Resexindex    mechanics.Rei
	TidallyLocked bool
}

func GeneratePlanet(orbit int, starType stars.StarType, s rand.Source) Planet {
	r := rand.New(s)

	orbitalRadius := mechanics.Radius{Radius: orbit, Unit: mechanics.LightMinute}

	a := r.Intn(99) + 1
	mass := getMass(a)

	zone, isLocked := getFormationZone(orbitalRadius.Radius, starType)

	planetType := getPlanetType(zone, mass)

	a = r.Intn(99) + 1
	m := generateMoons(planetType, mass, a)

	index := 0
	if planetType == TypeT || planetType == TypeSt {
		index = r.Intn(9) + 1
	}

	return Planet{
		Mass:          mass,
		Moons:         m,
		Orbit:         orbitalRadius,
		PlanetType:    planetType,
		HIndex:        index,
		Resexindex:    mechanics.VeryPoor,
		TidallyLocked: isLocked,
	}
}

func getMass(randInt int) PlanetMass {
	if randInt >= 3 && randInt < 6 {
		return Na
	}
	if randInt >= 6 && randInt < 26 {
		return MassOne
	}
	if randInt >= 26 && randInt < 76 {
		return MassTwo
	}
	if randInt >= 76 {
		return MassThree
	}
	return None
}

func getFormationZone(orbit int, starType stars.StarType) (PlanetFormationZone, bool) {
	isTideLocked := false
	var zone PlanetFormationZone
	switch starType {
	case stars.WhiteStar:
		{
			if orbit <= 5 {
				isTideLocked = true
			}
			if orbit < 20 {
				zone = Hot
			}
			if orbit >= 20 && orbit <= 40 {
				zone = Lwz
			}
			if orbit > 40 && orbit < 51 {
				zone = Cold
			}
			if orbit >= 51 && orbit <= 280 {
				zone = Gas
			}
			if orbit > 280 {
				zone = Ice
			}
		}
	case stars.YellowWhiteStar:
		{
			if orbit <= 4 {
				isTideLocked = true
			}
			if orbit < 10 {
				zone = Hot
			}
			if orbit >= 10 && orbit <= 18 {
				zone = Lwz
			}
			if orbit > 18 && orbit < 26 {
				zone = Cold
			}
			if orbit >= 26 && orbit <= 130 {
				zone = Gas
			}
			if orbit > 130 {
				zone = Ice
			}
		}
	case stars.YellowStar:
		{
			if orbit <= 3 {
				isTideLocked = true
			}
			if orbit < 6 {
				zone = Hot
			}
			if orbit >= 6 && orbit <= 12 {
				zone = Lwz
			}
			if orbit > 12 && orbit < 17 {
				zone = Cold
			}
			if orbit >= 17 && orbit <= 83 {
				zone = Gas
			}
			if orbit > 83 {
				zone = Ice
			}
		}
	case stars.OrangeStar:
		{
			if orbit <= 2 {
				isTideLocked = true
			}
			if orbit < 3 {
				zone = Hot
			}
			if orbit >= 3 && orbit <= 5 {
				zone = Lwz
			}
			if orbit > 5 && orbit < 10 {
				zone = Cold
			}
			if orbit >= 10 && orbit <= 38 {
				zone = Gas
			}
			if orbit > 38 {
				zone = Ice
			}
		}
	case stars.RedStar:
		{
			if orbit <= 2 {
				isTideLocked = true
			}
			if orbit < 3 {
				zone = Hot
			}
			if orbit >= 3 && orbit <= 4 {
				zone = Lwz
			}
			if orbit > 4 && orbit < 6 {
				zone = Cold
			}
			if orbit >= 6 && orbit <= 18 {
				zone = Gas
			}
			if orbit > 18 {
				zone = Ice
			}
		}
	case stars.RedDwarf:
		{
			if orbit == 1 {
				isTideLocked = true
			}
			if orbit < 4 {
				zone = Cold
			}
			if orbit >= 4 && orbit <= 11 {
				zone = Gas
			}
			if orbit > 11 {
				zone = Ice
			}
		}
	}
	return zone, isTideLocked
}

func getPlanetType(zone PlanetFormationZone, mass PlanetMass) PlanetType {
	switch zone {
	case Hot:
		{
			switch mass {
			case Na:
				{
					return AstB
				}
			case MassOne:
				{
					return TypeH
				}
			case MassTwo, MassThree:
				{
					return TypeV
				}
			}
		}
	case Lwz:
		{
			switch mass {
			case None:
			case Na:
				{
					return AstB
				}
			case MassOne:
				{
					return TypeB
				}
			case MassTwo:
				{
					return TypeT
				}
			case MassThree:
				{
					return TypeSt
				}
			}

		}
	case Cold:
		{
			switch mass {
			case Na:
				{
					return AstB
				}
			case MassOne, MassTwo, MassThree:
				{
					return TypeB
				}
			}

		}
	case Gas:
		{
			switch mass {
			case Na:
				{
					return AstB
				}
			case MassOne:
				{
					return TypeB
				}
			case MassTwo, MassThree:
				{
					return TypeG
				}
			}

		}
	case Ice:
		{
			switch mass {
			case Na:
				{
					return AstF
				}
			case MassOne:
				{
					return TypeF
				}
			case MassTwo, MassThree:
				{
					return TypeI
				}
			}

		}
	}
	return No
}

func generateMoons(planetType PlanetType, mass PlanetMass, num int) []moons.Moon {
	numMoons := 0
	switch planetType {
	case TypeB, TypeT, TypeSt:
		{
			switch mass {
			case MassOne:
				{
					numMoons = getNumberOfMoons(num - 50)
				}
			case MassTwo:
				{
					numMoons = getNumberOfMoons(num - 10)
				}
			case MassThree:
				{
					numMoons = getNumberOfMoons(num)
				}
			}
		}
	case TypeH:
		{
			switch mass {
			case MassOne:
				{
					numMoons = getNumberOfMoons(num - 65)
				}
			case MassTwo:
				{
					numMoons = getNumberOfMoons(num - 25)
				}
			case MassThree:
				{
					numMoons = getNumberOfMoons(num - 15)
				}
			}
		}
	case TypeF:
		{
			numMoons = getNumberOfMoons(num - 15)
		}
	case TypeV:
		{
			numMoons = getNumberOfMoons(num - 35)
		}
	case TypeI:
		{
			numMoons = getNumberOfMoons(num + 35)
		}
	case TypeG:
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

func generateMoon(isBig, isLocked bool, planetType PlanetType, reiNum, num, priorDistance int) moons.Moon {
	return moons.Moon{
		MType: getMoonType(planetType),
		MOrbit: mechanics.Radius{
			Radius: getMoonDistance(priorDistance, num, planetType),
			Unit:   mechanics.TacticalHex},
		IsBig:        isBig,
		IsTideLocked: isLocked,
		ResExpIndex:  getReiNumber(planetType, reiNum),
	}
}

func getMoonType(planetType PlanetType) moons.MoonType {
	switch planetType {
	case TypeB, TypeT, TypeSt, TypeG:
		{
			return moons.MB
		}
	case TypeH, TypeV:
		{
			return moons.MH
		}
	case TypeI, TypeF:
		{
			return moons.MF
		}
	}
	return moons.MB
}

func getMoonDistance(priorDistance, randNum int, planetType PlanetType) int {
	switch planetType {
	case TypeB, TypeF, TypeH, TypeSt, TypeT, TypeV:
		{
			priorDistance = mechanics.FractionRoundUp(priorDistance+randNum, 3)
		}
	case TypeG, TypeI:
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
func getReiNumber(planetType PlanetType, randNum int) mechanics.Rei {
	switch planetType {
	case TypeT, TypeSt, TypeV:
		{
			return determineRei(randNum)
		}
	case TypeG:
		{
			return determineRei(randNum + 1)
		}
	case TypeB:
		{
			return determineRei(randNum + 3)
		}
	case TypeF, TypeI, TypeH:
		{
			return determineRei(randNum + 5)
		}
	case AstF:
		{
			return mechanics.Normal
		}
	case AstB:
		{
			return mechanics.Rich
		}
	}
	return mechanics.VeryRich
}

func determineRei(randNum int) mechanics.Rei {
	if randNum < 2 {
		return mechanics.VeryPoor
	}
	if randNum >= 2 && randNum < 4 {
		return mechanics.Poor
	}
	if randNum >= 4 && randNum < 8 {
		return mechanics.Normal
	}
	if randNum >= 8 && randNum < 10 {
		return mechanics.Rich
	} else {
		return mechanics.VeryRich
	}
}

func modifyTideLockedPlanets() {

}
