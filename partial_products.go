package interview

// PartialProducts returns the set of partial products from a sorted set.
//
// Given a list of integers [1, 2, 3, 4], produce a set of partial
// products such that each index contains the product of the set, sans the current index.
//
// **Restrictions**
// 1. You cannot use division
// 2. Solution must be O(n)
//
// **Example**
// Input:  [1, 2, 3, 4]
// Output: [24, 12, 8, 6]
func PartialProducts(in []int) []int {

	// We'll solve this in two passes.
	//
	// 1. Starting from the beginning of the set,
	// we'll iterate through the values, storing
	// the previous product multiplied by the
	// previous input value.
	//
	// 2. Starting from the end of the set,
	// we'll iterate through the values, storing
	// the next product multiplied by the next
	// input value.
	//
	// Then we'll iterate over both the pre and
	// post sets, multiplying the values at each
	// index, storing the values in our final set.
	//
	// Example:
	// Input: [1, 2, 3, 4], Output: [24, 12, 8, 6]
	// Pre: [1, 1, 2, 6], Post: [24, 12, 4, 1]
	//
	// So let's look at index == 2
	// Pre[2]    = Pre[1]  * Input[1] == 1 * 2 == 2
	// Post[2]   = Post[3] * Input[3] == 1 * 4 == 4
	// Output[2] = Post[2] * Pre[2]   == 4 * 2 == 8

	pre := make([]int, len(in))
	for i := 0; i < len(in); i++ {

		// Our initial condition is 1.
		if i == 0 {

			pre[i] = 1
			continue
		}

		pre[i] = pre[i-1] * in[i-1]
	}

	post := make([]int, len(in))
	for i := len(in) - 1; i >= 0; i-- {

		// Same as above, our initial condition is 1.
		if i == len(in)-1 {

			post[i] = 1
			continue
		}

		post[i] = post[i+1] * in[i+1]
	}

	output := make([]int, len(in))
	for i, preVal := range pre {

		output[i] = preVal * post[i]
	}

	return output
}
