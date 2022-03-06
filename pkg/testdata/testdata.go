package testdata

import (
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/mechanics"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/stars"
)

var BGStar stars.Star = stars.Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{1, mechanics.LightMinute},
	},
	PlanetOrbitalLimit: stars.ZeroPlanetRadius,
	StarRadius:         stars.BlueGiantRadius,
	StarType:           stars.BlueGiant,
}

var RGStar stars.Star = stars.Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{1, mechanics.LightMinute},
	},
	PlanetOrbitalLimit: stars.ZeroPlanetRadius,
	StarRadius:         stars.RedGiantRadius,
	StarType:           stars.RedGiant,
}

var WStar stars.Star = stars.Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{1, mechanics.LightMinute},
	},
	PlanetOrbitalLimit: stars.YellowPlanetRadius,
	StarRadius:         stars.WhiteRadius,
	StarType:           stars.WhiteStar,
}

var YWStar stars.Star = stars.Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{1, mechanics.LightMinute},
	},
	PlanetOrbitalLimit: stars.YellowWhiteRadius,
	StarRadius:         stars.YellowWhiteRadius,
	StarType:           stars.YellowWhiteStar,
}

var YStar stars.Star = stars.Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{1, mechanics.LightMinute},
	},
	PlanetOrbitalLimit: stars.YellowWhiteRadius,
	StarRadius:         stars.YellowRadius,
	StarType:           stars.YellowStar,
}

var OStar stars.Star = stars.Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{1, mechanics.LightMinute},
	},
	PlanetOrbitalLimit: stars.OrangePlanetRadius,
	StarRadius:         stars.OrangeRadius,
	StarType:           stars.OrangeStar,
}

var RStar stars.Star = stars.Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{1, mechanics.LightMinute},
	},
	PlanetOrbitalLimit: stars.RedPlanetRadius,
	StarRadius:         stars.RedRadius,
	StarType:           stars.RedStar,
}

var RDStar stars.Star = stars.Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{1, mechanics.LightMinute},
	},
	PlanetOrbitalLimit: stars.RedPlanetRadius,
	StarRadius:         stars.RedDwarfRadius,
	StarType:           stars.RedDwarf,
}

var WDStar stars.Star = stars.Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{1, mechanics.LightMinute},
	},
	PlanetOrbitalLimit: stars.ZeroPlanetRadius,
	StarRadius:         stars.WhiteDwarfRadius,
	StarType:           stars.WhiteDwarf,
}

var NStar stars.Star = stars.Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{1, mechanics.LightMinute},
	},
	PlanetOrbitalLimit: stars.ZeroPlanetRadius,
	StarRadius:         stars.NeutronRadius,
	StarType:           stars.NeutronStar,
}

var Stars []stars.Star = []stars.Star{BGStar, RGStar, WStar, YWStar, YStar, OStar, RStar, RDStar, WDStar, NStar}
