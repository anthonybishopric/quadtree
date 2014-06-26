package quadtree

import (
// "fmt"
// "math"
// "sort"
)

const EARTH_RADIUS_MILES = 3959

const INTEGER_SCALAR = 1000000
const MIN_CELL_SIDE = 100

type GeoPlane struct {
	root *GeoNode
}

func InitPlane() *GeoPlane {
	return &GeoPlane{
		root: &GeoNode{latLow: 0, latHigh: 90, lonLow: 0, lonHigh: 180},
	}
}

type GeoNode struct {
	latLow      ScaledLat
	latHigh     ScaledLat
	lonLow      ScaledLon
	lonHigh     ScaledLon
	bottomLeft  *GeoNode
	bottomRight *GeoNode
	topLeft     *GeoNode
	topRight    *GeoNode
	locations   []Location
	minimum     bool
}

func (g *GeoNode) InRange(l *Location) bool {
	givenLat, givenLon := l.Scaled()
	inLat := g.latLow <= givenLat && g.latHigh > givenLat
	inLon := g.lonLow <= givenLon && g.lonHigh > givenLon
	return inLat && inLon
}

func (g *GeoNode) AddSubNode(node *GeoNode, lat ScaledLat, lon ScaledLon) {
	// highestBit := 1 << 28
}

// returns the smallest possible square on the Grid
// that contains both nodes
func (gp *GeoPlane) minSharedSquare(l, r GeoNode) *GeoNode {

	return nil
}

// fills a sparse quadtree with
func (gp *GeoPlane) addLocation(l *Location) *GeoNode {
	// scaledLat, scaledLon := l.Scaled()
	// leafNode := &GeoNode{
	// 	(scaledLat / MIN_CELL_SIDE) * MIN_CELL_SIDE,
	// 	(scaledLat + 1/MIN_CELL_SIDE) * MIN_CELL_SIDE,
	// 	(scaledLon / MIN_CELL_SIDE) * MIN_CELL_SIDE,
	// 	(scaledLon + 1/MIN_CELL_SIDE) * MIN_CELL_SIDE,
	// }
	// leafNode.addLocation(l)
	// return gp.root.AddSubNode(leafNode, scaledLat, scaledLon)
	return nil
}
