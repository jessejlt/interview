package interview

// NumericCompliment returns all numbers in the supplied list that
// contain both a positive and negative number.
func NumericCompliment(samples []int) []int {

	var (
		compliments []int
		sets        = make(map[int]int)
	)

	for _, sample := range samples {

		posSample := intAbs(sample)
		if val, ok := sets[posSample]; ok && val+sample == 0 {

			compliments = append(compliments, posSample)
			sets[posSample] *= 2
		} else if !ok {

			sets[posSample] = sample
		}
	}

	return compliments
}
