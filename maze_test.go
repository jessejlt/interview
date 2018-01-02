package interview_test

import (
	"testing"

	"github.com/jessejlt/interview"
)

func TestMinimumStepsThroughMaze(t *testing.T) {

	tests := []struct {
		Maze        [][]bool
		StartRow    int
		StartColumn int
		EndRow      int
		EndColumn   int
		Expect      int
		Name        string
	}{
		{
			Maze: [][]bool{
				[]bool{false, false, false, false},
				[]bool{true, true, false, true},
				[]bool{false, false, false, false},
				[]bool{false, false, false, false},
			},
			StartRow:    3,
			StartColumn: 0,
			EndRow:      0,
			EndColumn:   0,
			Expect:      7,
			Name:        "Example",
		},
	}

	for _, tt := range tests {

		have := interview.MinimumStepsThroughMaze(
			tt.Maze,
			tt.StartRow,
			tt.StartColumn,
			tt.EndRow,
			tt.EndColumn,
		)
		if have != tt.Expect {
			t.Fatalf("Have=%d, Expect=%d", have, tt.Expect)
		}
	}
}
