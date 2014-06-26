package quadtree

import (
	// "fmt"
	. "github.com/anthonybishopric/gotcha"
	"sort"
	"testing"
)

func TestSortsLocations(t *testing.T) {
	locations := []Location{
		Location{Lat: 39.908309, Lon: 116.418187},
		Location{Lat: 51.151176, Lon: 71.456451},
		Location{Lat: 37.760895, Lon: -122.434993},
		Location{Lat: 51.165549, Lon: 71.427376},
		Location{Lat: 37.760830, Lon: -122.436099},
		Location{Lat: 39.920159, Lon: 116.396858},
	}

	sort.Sort(ByZOrder(locations))

	expected := []Location{
		Location{37.760830, -122.436099},
		Location{37.760895, -122.434993},
		Location{51.151176, 71.456451},
		Location{51.165549, 71.427376},
		Location{39.908309, 116.418187},
		Location{39.920159, 116.396858},
	}

	debug := Locations(locations).Debug()
	for i, a := range locations {
		e := expected[i]
		Assert(t).AreEqual(a, e, debug)
	}
}
