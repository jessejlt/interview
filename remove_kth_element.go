package interview

// RemoveKthElement removes the kth last element in the list.
func RemoveKthElement(head *ListElement, k int) {

	if head == nil {
		return
	}

	var (
		previous = head
		current  = head
	)
	for head != nil {

		if k == 0 {

			previous = current
			current = current.Next
		}

		head = head.Next

		if k != 0 {
			k--
		}
	}

	if k == 0 {

		previous.Next = current.Next
	}
}
