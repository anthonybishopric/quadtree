package quadtree

import (
	. "github.com/anthonybishopric/gotcha"
	"testing"
)

func TestRootOnlyTree(t *testing.T) {
	root := &GeoNode{
		latLow:  0,
		latHigh: 100,
		lonLow:  0,
		lonHigh: 100,
	}

	Assert(t).AreEqual(true, root.InRange(&Location{50, 50}), "50,50 should be in (0--100,0--100) range")
}
