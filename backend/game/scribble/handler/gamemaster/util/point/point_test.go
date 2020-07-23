package point

import (
	"testing"
)

var TestGetPointsProvider = []struct {
	name           string
	expectedPoints int
}{
	{
		name:           "first iteration",
		expectedPoints: 50,
	},
	{
		name:           "second iteration",
		expectedPoints: 20,
	},
	{
		name:           "third iteration",
		expectedPoints: 10,
	},
	{
		name:           "fourth iteration",
		expectedPoints: 10,
	},
}

func TestGetPoints(t *testing.T) {
	pointHandler := Get()
	for _, testVal := range TestGetPointsProvider {
		t.Run(testVal.name, func(t *testing.T) {
			points := pointHandler.GetPoints()
			if testVal.expectedPoints != points {
				t.Errorf("expected: %v, got: %v", testVal.expectedPoints, points)
			}
		})
	}
}

func TestResetPoints(t *testing.T) {
	expectedPoints := 50

	pointHandler := Get()
	_ = pointHandler.GetPoints()
	_ = pointHandler.GetPoints()
	pointHandler.ResetPoints()
	points := pointHandler.GetPoints()
	if expectedPoints != points {
		t.Errorf("expected: %v, got: %v", expectedPoints, points)
	}

}
