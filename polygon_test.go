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
		{
			Polygon: Polygon{
				{
					1.3115500260580089,
					103.9826774597168,
				},
				{
					1.306916378856047,
					104.03383255004881,
				},
				{
					1.3566847287158743,
					104.04670715332031,
				},
				{
					1.3783079018157476,
					104.00842666625977,
				},
				{
					1.39443934768891,
					103.99538040161131,
				},
				{
					1.3467311394464763,
					103.97066116333008,
				},
				{
					1.333001984156086,
					103.97203445434569,
				},
				{
					1.3115500260580089,
					103.9826774597168,
				},
			},
			Point: Point{
				// Tightly within the boundary
				1.354868,
				103.990506,
			},
			Expect: true,
		},
		{
			Polygon: Polygon{
				{53.42498336008592, -6.245053394714894},
				{53.42503296303261, -6.247773493699079},
				{53.4238866509944, -6.263444610218625},
				{53.42083068230268, -6.264776931942881},
				{53.42069774635554, -6.261106657842928},
				{53.42005804788969, -6.253782000620896},
				{53.41994162898452, -6.248894885194622},
				{53.42060876079488, -6.248714072663663},
				{53.42167193520408, -6.246623133511962},
				{53.42150672283812, -6.242311594099218},
				{53.42200974116719, -6.242089191986312},
				{53.42229308506872, -6.239586938506013},
				{53.42177070342287, -6.237750455607769},
				{53.42243514235041, -6.237285912818003},
				{53.42323456462552, -6.23929527939592},
				{53.42418798796344, -6.238466912960044},
				{53.42542429652301, -6.237473530817869},
				{53.42564777901597, -6.236664314317801},
				{53.42603131387584, -6.236513248700376},
				{53.42684071147882, -6.238874212801099},
				{53.427193049176, -6.238722991835269},
				{53.42890374703639, -6.238075143654318},
				{53.42898650716833, -6.239146541111711},
				{53.42733353657361, -6.241069113221714},
				{53.42671604712674, -6.242069512459693},
				{53.42498336008592, -6.245053394714894},
			},
			Point: Point{
				53.426428,
				-6.238812,
			},
			Expect: true,
		},
	}

	// Iterate test cases and execute
	for _, test := range testCases[len(testCases)-1:] {

		res := test.Polygon.PointInside(test.Point)

		if res != test.Expect {

			// Report error
			t.Errorf("Expected point test to be %t, was %t", test.Expect, res)
		}
	}
}
