package interview_test

import (
	"testing"

	"github.com/jessejlt/interview"
)

func TestInterleaveHalfs(t *testing.T) {

	tests := []struct {
		Input  *interview.Stack
		Expect []int
		Name   string
	}{
		{
			Input:  interview.NewStack([]int{5, 4, 3, 2, 1}),
			Expect: []int{1, 5, 2, 4, 3},
			Name:   "Example",
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {

			have := interview.InterleaveHalfs(tt.Input)
			if !equalIntSlicesNoSort(have, tt.Expect) {

				t.Fatalf("Expect=%v, Have=%v", tt.Expect, have)
			}
		})
	}
}
