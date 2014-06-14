package quadtree

import (
	"bytes"
	"fmt"
)

type Location struct {
	Lat float64
	Lon float64
}

func (l Location) String() string {
	return fmt.Sprintf("(%.6f,%.6f)", l.Lat, l.Lon)
}

type Locations []Location

func (ls Locations) Debug() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, l := range ls {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(l.String())
	}
	buffer.WriteString("]")
	return buffer.String()
}
