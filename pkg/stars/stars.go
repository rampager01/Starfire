package stars

import (
	"math"
	"math/rand"
	"sort"

	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/mechanics"
)

type StellarRadius mechanics.Radius
type StarType string
type PlanetaryOrbitLimit int

const (
	BlueGiant       StarType = "Blue Giant"
	WhiteStar       StarType = "White Star"
	YellowWhiteStar StarType = "Yellow-White Star"
	YellowStar      StarType = "Yellow Star"
	OrangeStar      StarType = "Orange Star"
	RedStar         StarType = "Red Star"
	RedDwarf        StarType = "Red Dwarf"
	WhiteDwarf      StarType = "White Dwarf"
	RedGiant        StarType = "Red Giant"
	NeutronStar     StarType = "Neutron Star"
)

var BlueGiantRadius = StellarRadius(mechanics.Radius{Radius: 2, Unit: mechanics.LightMinute})
var WhiteRadius = StellarRadius(mechanics.Radius{Radius: 20, Unit: mechanics.TacticalHex})
var YellowWhiteRadius = StellarRadius(mechanics.Radius{Radius: 14, Unit: mechanics.TacticalHex})
var YellowRadius = StellarRadius(mechanics.Radius{Radius: 12, Unit: mechanics.TacticalHex})
var OrangeRadius = StellarRadius(mechanics.Radius{Radius: 8, Unit: mechanics.TacticalHex})
var RedRadius = StellarRadius(mechanics.Radius{Radius: 4, Unit: mechanics.TacticalHex})
var RedDwarfRadius = StellarRadius(mechanics.Radius{Radius: 2, Unit: mechanics.TacticalHex})
var WhiteDwarfRadius = StellarRadius(mechanics.Radius{Radius: 0, Unit: mechanics.TacticalHex})
var RedGiantRadius = StellarRadius(mechanics.Radius{Radius: 6, Unit: mechanics.LightMinute})
var NeutronRadius = StellarRadius(mechanics.Radius{Radius: 0, Unit: mechanics.TacticalHex})

const (
	ZeroPlanetRadius   PlanetaryOrbitLimit = 0
	RedPlanetRadius    PlanetaryOrbitLimit = 200
	OrangePlanetRadius PlanetaryOrbitLimit = 250
	YellowPlanetRadius PlanetaryOrbitLimit = 300
)

type Star struct {
	Orbit              mechanics.Orbit
	PlanetOrbitalLimit PlanetaryOrbitLimit
	StarRadius         StellarRadius
	StarType           StarType
}

//CalculateTitusBode calculates the Titus Bode planetary distance relationship.  This function provides the orbial radii of the planets orbiting a given star
func CalculateTitusBode(star Star, s rand.Source) []int {
	diffTest := false
	orbits := []int{}
	atLimit := false

	if star.PlanetOrbitalLimit == 0 {
		atLimit = true
		return orbits
	}

	for !diffTest {
		//		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		a := r.Intn(9) + 1
		b := r.Intn(9) + 1

		if math.Abs(float64(a-b)) >= 2 {
			orbits = append(orbits, a, b)
			sort.Ints(orbits)
			diffTest = true
		}
	}

	n := 1

	for !atLimit {
		orbit := orbits[0] + ((orbits[1] - orbits[0]) * (int(math.Pow(2, float64(n)))))

		if orbit <= int(star.PlanetOrbitalLimit) {
			orbits = append(orbits, orbit)
			n++
		} else {
			atLimit = true
		}
	}
	return orbits
}
