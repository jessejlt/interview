package interview_test

import (
	"testing"

	"github.com/jessejlt/interview"
)

func TestFindMaxRainVolume(t *testing.T) {

	tests := []struct {
		Topology []int
		Expect   int
		Name     string
	}{
		{
			Topology: []int{1, 3, 5, 1, 7, 2, 1},
			Expect:   4,
			Name:     "Single Dip",
		},
		{
			Topology: []int{1, 3, 5, 1, 7, 2, 3, 6},
			Expect:   11,
			Name:     "Double Dip",
		},
		{
			Topology: []int{1, 2, 3, 4},
			Expect:   0,
			Name:     "Ascending",
		},
		{
			Topology: []int{4, 3, 2, 1},
			Expect:   0,
			Name:     "Descending",
		},
		{
			Topology: []int{4, 3},
			Expect:   0,
			Name:     "No Valley",
		},
		{
			Topology: []int{},
			Expect:   0,
			Name:     "Empty",
		},
		{
			Topology: []int{2, 1, 2},
			Expect:   1,
			Name:     "Small Set",
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {

			have := interview.FindMaxRainVolume(tt.Topology)
			if have != tt.Expect {
				t.Fatalf("Expect=%d, Have=%d", tt.Expect, have)
			}
		})
	}
}
