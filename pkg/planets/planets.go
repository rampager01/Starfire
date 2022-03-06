package planets

import (
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
	PType         PlanetType
	HIndex        int
	Resexindex    mechanics.Rei
	TidallyLocked bool
}

func GetMass(randInt int) PlanetMass {
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

func GetFormationZone(orbit int, starType stars.StarType) (PlanetFormationZone, bool) {
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

func GetPlanetType(zone PlanetFormationZone, mass PlanetMass) PlanetType {
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

func GetReiNumber(planetType PlanetType, randNum int) mechanics.Rei {
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
