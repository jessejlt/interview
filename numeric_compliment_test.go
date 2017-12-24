package interview_test

import (
	"testing"

	"github.com/jessejlt/interview"
)

func TestNumericCompliment(t *testing.T) {

	tests := []struct {
		Samples []int
		Expect  []int
		Name    string
	}{
		{
			Samples: []int{2, 1, 3, -2, -1},
			Expect:  []int{2, 1},
			Name:    "Example",
		},
		{
			Samples: []int{},
			Expect:  []int{},
			Name:    "No Samples",
		},
		{
			Samples: []int{-2, -2, 2, 2},
			Expect:  []int{2},
			Name:    "Duplicates",
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {

			have := interview.NumericCompliment(tt.Samples)
			if !equalIntSlices(tt.Expect, have) {

				t.Fatalf("Expect=%v, Have=%v", tt.Expect, have)
			}
		})
	}
}
