package quadtree

import (
	"math"
)

// sort.Sort(ByZOrder(nodes))
type ByZOrder []Location

func (a ByZOrder) Len() int {
	return len(a)
}

func (a ByZOrder) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByZOrder) Less(i, j int) bool {
	// Implement sort described in http://en.wikipedia.org/wiki/Z-order_(curve)#Efficiently_building_quadtrees
	left, right := a[i], a[j]

	leftLat := a.scaleLat(left.Lat)
	rightLat := a.scaleLat(right.Lat)

	// the most significant 1 in latXor is the highest bit that is different between the two lats
	latXor := leftLat ^ rightLat

	leftLon := a.scaleLon(left.Lon)
	rightLon := a.scaleLon(right.Lon)

	lonXor := leftLon ^ rightLon

	// test whether lat or lon has largest different most significant bit
	lonIsBigger := latXor < lonXor && latXor < (latXor^lonXor)

	if lonIsBigger {
		return (leftLon - rightLon) < 0
	} else {
		return (leftLat - rightLat) < 0
	}
}

// these functions integerize the lat/long space,
// allowing us to order GeoNodes into a Z-order curve
// using a simpler bitwise comparison.
func (a ByZOrder) scaleLat(lat float64) int {
	return int(((lat + 90) * math.Pow10(6)))
}

func (a ByZOrder) scaleLon(lon float64) int {
	return int(((lon + 180) * math.Pow10(6)))
}
