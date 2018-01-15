package interview

// InterleaveHalfs returns a slice containing the first half of the values in
// the supplied stack with the second half of the values in reverse.
func InterleaveHalfs(values *Stack) []int {

	// First we'll move all our values out of the
	// stack and into a slice, thereby leaving
	// our stack and queue as buffers.
	var output []int
	for !values.IsEmpty() {

		output = append(
			output,
			values.Pop().(int),
		)
	}

	// Now iterate through our list, adding the first
	// half to our queue, and the second half to our stack.
	q := NewQueue()
	middle := len(output) / 2
	for i := 0; i < middle; i++ {

		q.Add(output[i])
	}

	for i := middle; i < len(output); i++ {

		values.Push(output[i])
	}

	// Finally, populate our final output slice.
	for i := 0; i < len(output); i++ {

		if !q.IsEmpty() {

			v := q.Remove().(int)
			output[i] = v
			i++
		}

		if !values.IsEmpty() {

			v := values.Pop().(int)
			output[i] = v
		}
	}

	return output
}
