package interview_test

import (
	"testing"

	"github.com/jessejlt/interview"
)

func TestQueue(t *testing.T) {

	tests := []struct {
		Input  []int
		Expect []int
		Name   string
	}{
		{
			Input:  []int{1, 2, 3},
			Expect: []int{1, 2, 3},
			Name:   "Example",
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {

			q := interview.NewQueue()
			for _, val := range tt.Input {

				q.Add(val)
			}

			for i := 0; !q.IsEmpty(); i++ {

				if have, expect := q.Remove().(int), tt.Input[i]; have != expect {

					t.Fatalf("Idx=%d, Have=%d, Expect=%d", i, have, expect)
				}
			}
		})
	}
}
