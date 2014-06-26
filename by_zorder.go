package quadtree

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

	leftLat, leftLon := left.Scaled()
	rightLat, rightLon := right.Scaled()

	// the most significant 1 in latXor is the highest bit that is different between the two lats
	latXor := int(leftLat ^ rightLat)
	lonXor := int(leftLon ^ rightLon)

	// test whether lat or lon has largest different most significant bit
	lonIsBigger := latXor < lonXor && latXor < (latXor^lonXor)

	if lonIsBigger {
		return (leftLon - rightLon) < 0
	} else {
		return (leftLat - rightLat) < 0
	}
}
