package main

import (
	"fmt"

	"github.com/ganners/gis-polygon"
)

func main() {

	polygon := polygon.Polygon{
		{
			51.8526099326355,
			-8.497022653516693,
		},
		{
			51.8441842566417,
			-8.49352118162505,
		},
		{
			51.84478200723391,
			-8.484808754096797,
		},
		{
			51.85351012088259,
			-8.488399836722362,
		},
	}

	point := polygon.Centroid()
	fmt.Printf("The centroid is %.6f %.6f\n", point.Lat, point.Lng)
}
