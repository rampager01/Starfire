package stars

import (
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/mechanics"
)

type StarData struct {
	Star      Star
	TitusBode []int
}

var BGStar Star = Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{Radius: 1, Unit: mechanics.LightMinute},
	},
	PlanetOrbitalLimit: ZeroPlanetRadius,
	StarRadius:         BlueGiantRadius,
	StarType:           BlueGiant,
}

var RGStar Star = Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{Radius: 1, Unit: mechanics.LightMinute},
	},
	PlanetOrbitalLimit: ZeroPlanetRadius,
	StarRadius:         RedGiantRadius,
	StarType:           RedGiant,
}

var WStar Star = Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{Radius: 1, Unit: mechanics.LightMinute},
	},
	PlanetOrbitalLimit: YellowPlanetRadius,
	StarRadius:         WhiteRadius,
	StarType:           WhiteStar,
}

var YWStar Star = Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{Radius: 1, Unit: mechanics.LightMinute},
	},
	PlanetOrbitalLimit: YellowPlanetRadius,
	StarRadius:         YellowWhiteRadius,
	StarType:           YellowWhiteStar,
}

var YStar Star = Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{Radius: 1, Unit: mechanics.LightMinute},
	},
	PlanetOrbitalLimit: YellowPlanetRadius,
	StarRadius:         YellowRadius,
	StarType:           YellowStar,
}

var OStar Star = Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{Radius: 1, Unit: mechanics.LightMinute},
	},
	PlanetOrbitalLimit: OrangePlanetRadius,
	StarRadius:         OrangeRadius,
	StarType:           OrangeStar,
}

var RStar Star = Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{Radius: 1, Unit: mechanics.LightMinute},
	},
	PlanetOrbitalLimit: RedPlanetRadius,
	StarRadius:         RedRadius,
	StarType:           RedStar,
}

var RDStar Star = Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{Radius: 1, Unit: mechanics.LightMinute},
	},
	PlanetOrbitalLimit: RedPlanetRadius,
	StarRadius:         RedDwarfRadius,
	StarType:           RedDwarf,
}

var WDStar Star = Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{Radius: 1, Unit: mechanics.LightMinute},
	},
	PlanetOrbitalLimit: ZeroPlanetRadius,
	StarRadius:         WhiteDwarfRadius,
	StarType:           WhiteDwarf,
}

var NStar Star = Star{
	Orbit: mechanics.Orbit{
		Bearing: 12,
		Radius:  mechanics.Radius{Radius: 1, Unit: mechanics.LightMinute},
	},
	PlanetOrbitalLimit: ZeroPlanetRadius,
	StarRadius:         NeutronRadius,
	StarType:           NeutronStar,
}

var Data []StarData = []StarData{
	{Star: BGStar, TitusBode: []int{}},
	{Star: RGStar, TitusBode: []int{}},
	{Star: WStar, TitusBode: []int{3, 5, 7, 11, 19, 35, 67, 131, 259}},
	{Star: YWStar, TitusBode: []int{3, 5, 7, 11, 19, 35, 67, 131, 259}},
	{Star: YStar, TitusBode: []int{3, 5, 7, 11, 19, 35, 67, 131, 259}},
	{Star: OStar, TitusBode: []int{3, 5, 7, 11, 19, 35, 67, 131}},
	{Star: RStar, TitusBode: []int{3, 5, 7, 11, 19, 35, 67, 131}},
	{Star: RDStar, TitusBode: []int{3, 5, 7, 11, 19, 35, 67, 131}},
	{Star: WDStar, TitusBode: []int{}},
	{Star: NStar, TitusBode: []int{}},
}

var Stars []Star = []Star{
	BGStar,
	RGStar,
	WStar,
	YWStar,
	YStar,
	OStar,
	RStar,
	RDStar,
	WDStar,
	NStar,
}
