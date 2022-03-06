package moons

import "github.com/rampager01/starfire-solar-galaxy-generator/pkg/mechanics"

type MoonType string

const (
	MB MoonType = "mB"
	MF MoonType = "mF"
	MH MoonType = "mH"
	MT MoonType = "mT"
)

type Moon struct {
	MType        MoonType
	MOrbit       mechanics.Radius
	IsBig        bool
	IsTideLocked bool
	ResExpIndex  mechanics.Rei
}

func CheckTideLocked(numMoons, rand int) (bool, bool) {
	isBig := false
	isTideLocked := false
	if rand <= 7+numMoons {
		isTideLocked = true
		if rand <= 1+numMoons {
			isBig = true
		}
	}
	return isTideLocked, isBig
}
