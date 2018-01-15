package interview_test

import (
	"testing"

	"github.com/jessejlt/interview"
)

func TestAreBracesWellFormed(t *testing.T) {

	tests := []struct {
		Input  string
		Expect bool
	}{
		{
			Input:  "([])[]({})",
			Expect: true,
		},
		{
			Input:  "([)]",
			Expect: false,
		},
		{
			Input:  "((()",
			Expect: false,
		},
		{
			Input:  "({([])})",
			Expect: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.Input, func(t *testing.T) {

			have := interview.AreBracesWellFormed(tt.Input)
			if have != tt.Expect {
				t.Fatalf("Have=%t, Expect=%t", have, tt.Expect)
			}
		})
	}
}
