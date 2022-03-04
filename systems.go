package system

type anomaly struct {
  
}

type starlessNexus struct {
  warpPoints []warpPoint
}

type starlessNebula struct {
  warpPoints []warpPoint
}

type singleStar struct {
  star: star
  warpPoints []warpPoint
}

type singleStarNebula struct {
  star star
  nebula nebula
  warpPoints []warpPoint
} 

type binaryStar struct {
  stars []star
  warpPoints []warpPoint
}

type binayStarNebula struct {
  stars []star
  nebula nebula
  warpPoints []warpPoint
}

type trinaryStar struct {
  stars []star
  warpPoints []warpPoint
}

type starSystem interface {
	info()
  generatePlanets()
  generateWarpPoints()
}
