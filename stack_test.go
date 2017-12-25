package interview_test

import (
	"testing"

	"github.com/jessejlt/interview"
)

func TestCtorStack(t *testing.T) {

	tests := []struct {
		Input  *interview.Stack
		Expect []int
		Name   string
	}{
		{
			Input:  interview.NewStack([]int{1, 2, 3}),
			Expect: []int{3, 2, 1},
			Name:   "Example",
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {

			for i := 0; !tt.Input.IsEmpty(); i++ {

				if have, expect := tt.Input.Pop(), tt.Expect[i]; have != expect {

					t.Fatalf("Idx=%d, Have=%d, Expect=%d", i, have, expect)
				}
			}
		})
	}
}

func TestStackPush(t *testing.T) {

	tests := []struct {
		Input  []int
		Expect []int
		Name   string
	}{
		{
			Input:  []int{1, 2, 3},
			Expect: []int{3, 2, 1},
			Name:   "Example",
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {

			s := interview.NewStack(nil)
			for _, val := range tt.Input {

				s.Push(val)
			}

			for i := 0; !s.IsEmpty(); i++ {

				if have, expect := s.Pop(), tt.Expect[i]; have != expect {

					t.Fatalf("Idx=%d, Have=%d, Expect=%d", i, have, expect)
				}
			}
		})
	}
}
