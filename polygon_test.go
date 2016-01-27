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
		{
			Polygon: Polygon{
				{2.048704121475433, 41.2869909760169},
				{2.052365319287515, 41.28347961914785},
				{2.065007219903774, 41.28281491400502},
				{2.072222631811531, 41.28542327523093},
				{2.070049770254769, 41.29011074828045},
				{2.057864904242872, 41.29640570768664},
				{2.078735038634543, 41.30404483220536},
				{2.092074164418243, 41.30962718101072},
				{2.078568301696546, 41.31867997986441},
				{2.076720185963297, 41.32212726252342},
				{2.07895976684928, 41.32600840606369},
				{2.068286506614412, 41.32319098216702},
				{2.069097934629491, 41.3180008590347},
				{2.061446890699019, 41.30701330329803},
				{2.046938375921601, 41.30289893326319},
				{2.049148577943003, 41.29551913362086},
				{2.049318341255011, 41.28735077253985},
				{2.048704121475433, 41.2869909760169},
			},
			Point: Point{
				// Tightly within the boundary
				2.053203,
				41.287651,
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
