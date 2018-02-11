package interview

// Permutations returns the permutations of a supplied list
func Permutations(in []int) [][]int {

	var permute func(in []int, l, r int)
	var permutations [][]int

	permute = func(in []int, l, r int) {

		if l == r {
			permutations = append(
				permutations,
				in,
			)
			return
		}

		for i := l; i <= r; i++ {

			in[l], in[i] = in[i], in[l]
			permute(in, l+1, r)
			in[l], in[i] = in[i], in[l]
		}
	}

	permute(in, 0, len(in)-1)
	return permutations
}
