package gotesting

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestComputeScore(t *testing.T) {


	testCases := []struct {
		Hand          string
		ExpectedScore int
	}{
		{"2", 2},
		{"3", 3},
		{"J", 10},
		{"A", 11},
		{"AJ", 21},
		{"A345", 23},
		{"/345≈", 0}, // Invalid character
		{"", 0},      // Empty hand
	}

	for _, entry := range testCases {
		Convey("With Hand "+entry.Hand, t, func() {
			So(ComputeScore(entry.Hand), ShouldEqual, entry.ExpectedScore)
		})
	}
}

