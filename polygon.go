package polygon

import "math"

type Point struct {
	Lat float64
	Lng float64
}

type Polygon []Point

// Area calculates the area of the polygon
func (p Polygon) Area() float64 {

	signedArea := 0.0
	for i := 0; i < len(p); i++ {
		iʹ := (i + 1) % (len(p))
		signedArea += (p[i].Lat * p[iʹ].Lng) - (p[i].Lng * p[iʹ].Lat)
	}
	signedArea = math.Abs(signedArea) / 2

	return signedArea
}

// Calculates the centroid for a given point, also returns the area
// which was calculated. Might be useful
func (p Polygon) Centroid() Point {

	signedArea := p.Area()

	x := 0.0
	y := 0.0
	for j := 0; j < len(p); j++ {
		jʹ := (j + 1) % (len(p))
		norm := (p[j].Lat * p[jʹ].Lng) - (p[jʹ].Lat * p[j].Lng)
		x += (p[j].Lat + p[jʹ].Lat) * norm
		y += (p[j].Lng + p[jʹ].Lng) * norm
	}

	x = math.Abs(x) * (1.0 / (6.0 * signedArea))
	y = math.Abs(y) * (1.0 / (6.0 * signedArea))

	return Point{
		Lat: x,
		Lng: y,
	}
}

// Determines if a point is inside or outside this polygon
func (p Polygon) PointInside(point Point) bool {

	numCollisions := 0

	for i := 0; i < len(p); i++ {

		iʹ := (i + 1) % (len(p))

		// If, with Lat, the point sits  within the edge
		if (p[i].Lat <= point.Lat) && (point.Lat < p[iʹ].Lat) ||
			(p[iʹ].Lat <= point.Lat) && (point.Lat < p[i].Lat) {

			dLng := math.Abs(p[iʹ].Lng - p[i].Lng)
			dLat := math.Abs(point.Lat - p[i].Lat)

			if point.Lng < ((dLng*dLat)/dLat)+p[i].Lng {

				// It's a collision, so iterate the numCollisions
				numCollisions++
			}
		}
	}

	// If the number of collisions is ODD, return true (it's inside)
	return ((numCollisions & 1) == 1)
}
