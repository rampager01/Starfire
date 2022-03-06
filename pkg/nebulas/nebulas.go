package nebulas

type NebulaType string
type NebulaSpecial string

const (
	Dark       NebulaType = "Dark Nebula"
	Emission   NebulaType = "Emission Nebula"
	Maser      NebulaType = "Maser Nebula"
	Reflective NebulaType = "Reflective Nebula"
)

const (
	Corrosive               NebulaSpecial = "Corrosive"
	DarkPulse               NebulaSpecial = "Dark Pulse"
	DatlinkUseless          NebulaSpecial = "Datalink Useless"
	DestabilizedDriveFields NebulaSpecial = "Destabilized Drive-Fields"
	EmissionFlare           NebulaSpecial = "Emission Flare"
	EnergyAbsorption        NebulaSpecial = "Energy Absorption"
	EnergySpike             NebulaSpecial = "Energy Spike"
	Ghosts                  NebulaSpecial = "Ghosts in the Mist"
	HgtPresent              NebulaSpecial = "HGT Present"
	LighningStrikes         NebulaSpecial = "Lightning Strikes/Energy Pulses"
	MistWraiths             NebulaSpecial = "Mist-Wraiths"
	NoSpecial               NebulaSpecial = "None"
	RadiationPulses         NebulaSpecial = "Radiation Pulses"
	Ragged                  NebulaSpecial = "Ragged"
	ReflectiveFlare         NebulaSpecial = "Reflective Flare"
	Sargassos               NebulaSpecial = "Sargassos"
	TargetingConfusion      NebulaSpecial = "Targeting Confusion"
	VariablePulseRate       NebulaSpecial = "Varaiable Pulse Rate"
)

type Nebula struct {
	NebulaType NebulaType
	Special    NebulaSpecial
}
