package mechanics

type DistanceUnit string

// can only be of value 1-12
type Radian int
type Rei string

const (
	TacticalHex DistanceUnit = "tH"
	LightMinute DistanceUnit = "LM"
	SystemHex   DistanceUnit = "sH"
	Stmp        DistanceUnit = "StMP"
)

const (
	ShToSmTP int = 30
	LmTosH   int = 12
	ThToLm   int = 240
)

const (
	VeryPoor  Rei = "Very Poor"
	Poor      Rei = "Poor"
	Normal    Rei = "Normal"
	Rich      Rei = "Rich"
	VeryRich  Rei = "Very Rich"
	SuperRich Rei = "Super Rich"
)

type Radius struct {
	Radius int
	Unit   DistanceUnit
}

type Orbit struct {
	Bearing Radian
	Radius  Radius
}

func FractionRoundUp(num, divisor int) int {
	if divisor == 2 {
		if num%2 == 0 {
			return num / divisor
		} else {
			return (num + 1) / divisor
		}
	}
	if divisor == 3 {
		if num%3 == 0 {
			return num / divisor
		}
		if num%3 == 1 {
			return (num + 2) / divisor
		}
		if num%3 == 2 {
			return (num + 1) / divisor
		}
	}
	return 0
}
