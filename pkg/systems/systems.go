package systems

import (
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/nebulas"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/planets"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/stars"
	"github.com/rampager01/starfire-solar-galaxy-generator/pkg/warppoints"
)

type SystemType string

const (
	AsteroidCluster SystemType = "Ateroid Cluster"
	BinaryNebula    SystemType = "Binary Star System"
	BinaryStar      SystemType = "Binary Star System"
	Blackhole       SystemType = "Blackhole"
	GravityWave     SystemType = "Gravity Wave Effect"
	LongDistance    SystemType = "Long-Distance Companion"
	Magnetar        SystemType = "Magnetar"
	ProtoDisk       SystemType = "Proto-Disk Formation"
	Pulsar          SystemType = "Pulsar"
	Quarternary     SystemType = "Quarternary Star System"
	SingleNebula    SystemType = "Single-Star Nebula"
	SingleStar      SystemType = "Single-Star Star System"
	StarlessNexus   SystemType = "Starless Nexus"
	StarlessNebula  SystemType = "Starless Nebula"
	TrinayStar      SystemType = "Trinary Star System"
)

type StarSystem struct {
	Nebula         []nebulas.Nebula
	Planets        []planets.Planet
	Stars          []stars.Star
	StarSystemType SystemType
	WarpPoints     []warppoints.WarpPoint
}
