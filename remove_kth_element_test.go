package interview_test

import (
	"testing"

	"github.com/jessejlt/interview"
)

func TestRemoveKthElement(t *testing.T) {

	tests := []struct {
		Head   *interview.ListElement
		K      int
		Expect []int
		Name   string
	}{
		{
			Head: interview.NewSinglyLinkedList(
				[]int{1, 2, 3, 4, 5, 6, 7},
			),
			K:      3,
			Expect: []int{1, 2, 3, 4, 6, 7},
			Name:   "Example",
		},
		{
			Head: interview.NewSinglyLinkedList(
				[]int{1, 2, 3},
			),
			K:      3,
			Expect: []int{1, 2, 3},
			Name:   "Too Small",
		},
	}

	for _, tt := range tests {

		interview.RemoveKthElement(tt.Head, tt.K)

		for i, expect := range tt.Expect {

			if tt.Head == nil {
				t.Fatalf("Idx=%d, Have=nil, Expect=%d", i, expect)
			}
			if tt.Head.Value != expect {
				t.Fatalf("Idx=%d, Have=%d, Expect=%d", i, tt.Head.Value, expect)
			}
			tt.Head = tt.Head.Next
		}
		if tt.Head != nil {
			t.Fatalf("Have trailing element=%d", tt.Head.Value)
		}
	}
}
