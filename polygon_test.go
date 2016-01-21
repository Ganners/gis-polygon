package polygon

import (
	"reflect"
	"testing"
)

// Tests the area
func TestArea(t *testing.T) {

	testCases := []struct {
		Polygon    Polygon
		ExpectArea float64
	}{
		{
			// Simple test with a square
			Polygon: Polygon{
				{1, 1},
				{3, 1},
				{3, 3},
				{1, 3},
			},
			ExpectArea: 4,
		},
		{
			// Simple test with a square
			Polygon: Polygon{
				{3, 4},
				{5, 11},
				{12, 8},
				{9, 5},
				{5, 6},
			},
			ExpectArea: 30,
		},
	}

	for _, test := range testCases {

		// Pull out the centroid
		area := test.Polygon.Area()

		// Check if the areas are equal
		if area != test.ExpectArea {
			t.Errorf("Area %f doesn't match expected %f", area, test.ExpectArea)
		}
	}
}

// Tests the centroid
func TestCentroid(t *testing.T) {

	testCases := []struct {
		Polygon     Polygon
		ExpectPoint Point
	}{
		{
			// Simple test with a square
			Polygon: Polygon{
				{1, 1},
				{3, 1},
				{3, 3},
				{1, 3},
			},
			ExpectPoint: Point{
				2.0,
				2.0,
			},
		},
		{
			// Simple test with a square
			Polygon: Polygon{
				{3, 4},
				{5, 11},
				{12, 8},
				{9, 5},
				{5, 6},
			},
			ExpectPoint: Point{
				7.166666666666667,
				7.611111111111112,
			},
		},
	}

	for _, test := range testCases {

		// Pull out the centroid
		res := test.Polygon.Centroid()

		// Check if the two structs are equal
		if !reflect.DeepEqual(res, test.ExpectPoint) {

			t.Errorf("Expected point to be %#v, was %#v", test.ExpectPoint, res)
		}
	}
}

// Tests the point is inside (only compatible with simple polygons)
func TestPointInside(t *testing.T) {

	testCases := []struct {
		Polygon Polygon
		Point   Point
		Expect  bool
	}{
		{
			Polygon: Polygon{
				{1, 1},
				{3, 1},
				{3, 3},
				{1, 3},
			},
			Point: Point{
				2.0,
				2.0,
			},
			Expect: true,
		},
		{
			Polygon: Polygon{
				{1, 1},
				{3, 1},
				{3, 3},
				{1, 3},
			},
			Point: Point{
				4.0,
				4.0,
			},
			Expect: false,
		},
		{
			Polygon: Polygon{
				{1, 1},
				{3, 1},
				{3, 3},
				{1, 3},
			},
			Point: Point{
				0.99999,
				3,
			},
			Expect: false,
		},
		{
			Polygon: Polygon{
				{1, 1},
				{3, 1},
				{3, 3},
				{1, 3},
			},
			Point: Point{
				// Tightly within the boundary
				1.00000001,
				2.99999999,
			},
			Expect: true,
		},
	}

	// Iterate test cases and execute
	for _, test := range testCases {

		res := test.Polygon.PointInside(test.Point)

		if res != test.Expect {

			// Report error
			t.Errorf("Expected point test to be %t, was %t", test.Expect, res)
		}
	}
}
