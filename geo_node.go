package quadtree

import (
	// "fmt"
	"math"
)

const EARTH_RADIUS_MILES = 3959

type GeoNode struct {
	latLow  float64
	latHigh float64
	lonLow  float64
	lonHigh float64
}

func (g *GeoNode) InRange(l *Location) bool {
	inLat := g.latLow <= l.Lat && g.latHigh > l.Lat
	inLon := g.lonLow <= l.Lon && g.lonHigh > l.Lon
	return inLat && inLon
}

// from http://mathforum.org/library/drmath/view/63767.html
func (g *GeoNode) Area() float64 {
	return (math.Pi / 180.0) * math.Pow(EARTH_RADIUS_MILES, 2) * math.Abs(math.Sin(g.latLow)-math.Sin(g.latHigh)) * math.Abs(g.lonLow-g.lonHigh)
}
