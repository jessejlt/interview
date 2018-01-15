package interview

import (
	"strings"
)

// AreBracesWellFormed returns true if the round, curly and square braces
// are balanced and well-formed.
func AreBracesWellFormed(in string) bool {

	s := NewStack()
	for _, brace := range strings.Split(in, "") {

		if isOpenBrace(brace) {
			s.Push(brace)
		} else {

			v, ok := s.Pop().(string)
			if !ok {
				return false
			}

			if !braceCloses(v, brace) {
				return false
			}
		}
	}

	return s.IsEmpty()
}

func isOpenBrace(brace string) bool {

	if brace == "(" ||
		brace == "[" ||
		brace == "{" {
		return true
	}
	return false
}

func braceCloses(a, b string) bool {

	if a == "(" && b == ")" {
		return true
	}
	if a == "[" && b == "]" {
		return true
	}
	if a == "{" && b == "}" {
		return true
	}

	return false
}
