package mechanics

type distanceUnit string
// can only be of value 1-12
type radian int

const (
  TacticalHex distanceUnit = "tH"
  LightMinute distanceUnit = "LM"
  SystemHex distanceUnit = "sH"
  Stmp distanceUnit = "StMP"
)

const (
  ShToSmTP int = 30
  LmTosH int = 12
  ThToLm int = 240
)

type radius struct {
  radius int
  unit distanceUnit
}

type orbit struct {
  bearing radian
  radius radius
}

