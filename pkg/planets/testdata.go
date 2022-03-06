package planets

import (
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/stars"
)

var testOrbits []int = []int{2, 5, 8, 14, 26, 50}

type result struct {
	zone       PlanetFormationZone
	tideLocked bool
}

type data struct {
	startype stars.StarType
	orbits   []int
	results  []result
}

var wStar data = data{
	startype: stars.WhiteStar,
	orbits:   testOrbits,
	results: []result{
		{Hot, true},
		{Hot, true},
		{Hot, false},
		{Hot, false},
		{Lwz, false},
		{Cold, false},
	},
}

var ywStar data = data{
	startype: stars.YellowWhiteStar,
	orbits:   testOrbits,
	results: []result{
		{Hot, true},
		{Hot, false},
		{Hot, false},
		{Lwz, false},
		{Gas, false},
		{Gas, false},
	},
}

var yStar data = data{
	startype: stars.YellowStar,
	orbits:   testOrbits,
	results: []result{
		{Hot, true},
		{Hot, false},
		{Lwz, false},
		{Cold, false},
		{Gas, false},
		{Gas, false},
	},
}

var oStar data = data{
	startype: stars.OrangeStar,
	orbits:   testOrbits,
	results: []result{
		{Hot, true},
		{Lwz, false},
		{Cold, false},
		{Gas, false},
		{Gas, false},
		{Ice, false},
	},
}

var rStar data = data{
	startype: stars.RedStar,
	orbits:   testOrbits,
	results: []result{
		{Hot, true},
		{Cold, false},
		{Gas, false},
		{Gas, false},
		{Ice, false},
		{Ice, false},
	},
}

var rdStar data = data{
	startype: stars.RedDwarf,
	orbits:   testOrbits,
	results: []result{
		{Cold, false},
		{Gas, false},
		{Gas, false},
		{Ice, false},
		{Ice, false},
		{Ice, false},
	},
}

var testData []data = []data{
	wStar,
	ywStar,
	yStar,
	oStar,
	rStar,
	rdStar,
}
