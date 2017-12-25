package interview_test

import (
	"sort"
)

func equalIntSlices(a, b []int) bool {

	sort.Ints(a)
	sort.Ints(b)

	return equalIntSlicesNoSort(a, b)
}

func equalIntSlicesNoSort(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	for i, v := range a {

		if b[i] != v {
			return false
		}
	}

	return true
}
