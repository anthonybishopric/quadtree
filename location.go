package quadtree

import (
	"bytes"
	"fmt"
	"math"
)

type Area float64

type Longitude float64
type Latitude float64

type ScaledLat int
type ScaledLon int

type Location struct {
	Lat Latitude
	Lon Longitude
}

// these functions integerize the lat/long space,
// allowing us to order GeoNodes into a Z-order curve
// using a simpler bitwise comparison.
func (lat Latitude) Scale() ScaledLat {
	return ScaledLat((float64(lat+90) * math.Pow10(6)))
}

func (lat Latitude) Float() float64 {
	return float64(lat)
}

func (lon Longitude) Scale() ScaledLon {
	return ScaledLon((float64(lon+180) * math.Pow10(6)))
}

func (lon Longitude) Float() float64 {
	return float64(lon)
}

func (l Location) Scaled() (ScaledLat, ScaledLon) {
	return l.Lat.Scale(), l.Lon.Scale()
}

func (l Location) String() string {
	return fmt.Sprintf("(%.6f,%.6f)", l.Lat, l.Lon)
}

func (l Location) Area(l2 Location) Area {
	return Area((math.Pi / 180.0) * math.Pow(EARTH_RADIUS_MILES, 2) * math.Abs(math.Sin(l.Lat.Float())-math.Sin(l2.Lat.Float())) * math.Abs(l.Lon.Float()-l2.Lon.Float()))
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
