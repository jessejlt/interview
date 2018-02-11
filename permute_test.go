package interview_test

import (
	"testing"

	"github.com/jessejlt/interview"
)

func TestPermuations(t *testing.T) {

	have := interview.Permutations([]int{1, 2, 3})
	if len(have) != 6 {
		t.Log(have)
		t.Fatalf("Have=%d, Expect=6", len(have))
	}
}
