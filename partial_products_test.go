package interview_test

import (
	"testing"

	"github.com/jessejlt/interview"
)

func TestPartialProducts(t *testing.T) {

	tests := []struct {
		In     []int
		Expect []int
		Name   string
	}{
		{
			In:     []int{1, 2, 3},
			Expect: []int{6, 3, 2},
			Name:   "Small Set",
		},
		{
			In:     []int{1, 2, 3, 4, 5},
			Expect: []int{120, 60, 40, 30, 24},
			Name:   "Large Set",
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {

			have := interview.PartialProducts(tt.In)

			if len(have) != len(tt.Expect) {

				t.Fatalf("Expect=%v, Have=%v", tt.Expect, have)
			}

			for i, v := range have {

				if tt.Expect[i] != v {

					t.Fatalf("Expect=%v, Have=%v", tt.Expect, have)
				}
			}
		})
	}
}
