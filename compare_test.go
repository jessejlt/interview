package interview_test

import (
	"sort"
)

func equalIntSlices(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i, v := range a {

		if b[i] != v {
			return false
		}
	}

	return true
}
