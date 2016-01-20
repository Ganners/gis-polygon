package collision

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

	for i, j := 0, (len(p) - 1); i < len(p); i++ {

		// If, with Lat, the point sits  within the edge
		if (p[i].Lat <= point.Lat) && (point.Lat < p[j].Lat) ||
			(p[j].Lat <= point.Lat) && (point.Lat < p[i].Lat) {

			// And it sits inside the edge in Lon
			dLng := (p[j].Lng - p[i].Lng)
			dLat := (point.Lat - p[i].Lat)

			if point.Lng < ((dLng*dLat)/dLat)+p[i].Lng {
				// if point.Lng < (p[j].Lng-p[i].Lng)*(point.Lat-p[i].Lat)/(p[j].Lat-p[i].Lat)+p[i].Lng {

				// It's a collision, so iterate the numCollisions
				numCollisions++
			}
		}

		j = i
	}

	// If the number of collisions is ODD, return true (it's inside)
	return ((numCollisions & 1) == 1)
}
