package interview

func intAbs(val int) int {

	if val < 0 {
		return val * -1
	}
	return val
}

func intMax(a, b int) int {

	if a < b {
		return b
	}
	return a
}
