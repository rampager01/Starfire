package warppoints

import "github.com/rampager01/starfire-solar-galaxy-generator/pkg/mechanics"

type WarpPointVisibility string
type WarpPointType string

const (
	Open        WarpPointVisibility = "Open"
	Concealed   WarpPointVisibility = "Concealed"
	Hidden      WarpPointVisibility = "Hidden"
	Secret      WarpPointVisibility = "Secret"
	TopSecret   WarpPointVisibility = "Top Secret"
	Undisclosed WarpPointVisibility = "Undisclosed"
)

const (
	A WarpPointType = "A"
	B WarpPointType = "B"
	C WarpPointType = "C"
	D WarpPointType = "D"
	E WarpPointType = "E"
	F WarpPointType = "F"
)

type WarpPoint struct {
	Orbit          mechanics.Orbit
	WPType         WarpPointType
	Visibility     WarpPointVisibility
	WPSystemLinkID int
}

func GetWarpPointType(wpType int) WarpPointType {
	if wpType == 1 {
		return A
	}
	if wpType >= 2 && wpType < 4 {
		return B
	}
	if wpType >= 4 && wpType < 7 {
		return C
	}
	if wpType >= 7 && wpType < 9 {
		return D
	}
	if wpType == 9 {
		return E
	} else {
		return F
	}
}

func GetWarpPointVisibility(visibility int) WarpPointVisibility {
	if visibility >= 1 && visibility < 7 {
		return Open
	}
	if visibility >= 7 && visibility < 9 {
		return Concealed
	}
	if visibility == 9 {
		return Hidden
	} else {
		return Secret
	}
}
