package interview_test

import (
	"testing"

	"github.com/jessejlt/interview"
)

func TestBuySellStocks(t *testing.T) {

	tests := []struct {
		Prices []int
		Expect int
		Name   string
	}{
		{
			Prices: []int{10, 12, 8, 14, 13},
			Expect: 6,
			Name:   "Example",
		},
		{
			Prices: []int{10, 5, 3, 1},
			Expect: 0,
			Name:   "Sell-Off",
		},
		{
			Prices: []int{1, 5, 3, 8},
			Expect: 7,
			Name:   "Max Bounds",
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {

			have := interview.BuySellStocks(tt.Prices)
			if tt.Expect != have {

				t.Fatalf("Expect=%d, Have=%d", tt.Expect, have)
			}
		})
	}
}
