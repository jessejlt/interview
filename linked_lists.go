package interview

// ListElement is an element from a linked list.
type ListElement struct {
	Value int
	Next  *ListElement
	Prev  *ListElement
}

// NewSinglyLinkedList returns the Head of a singly linked list
// populated with the specified values.
func NewSinglyLinkedList(values []int) *ListElement {

	if len(values) == 0 {
		return nil
	}

	head := &ListElement{Value: values[0]}
	ptr := head
	for i := 1; i < len(values); i++ {

		ptr.Next = &ListElement{Value: values[i]}
		ptr = ptr.Next
	}

	return head
}

// NewDoublyLinkedList returns the Head of a doubly linked list
// populated with the specified values.
func NewDoublyLinkedList(values []int) *ListElement {

	if len(values) == 0 {
		return nil
	}

	head := &ListElement{Value: values[0]}
	ptr := head
	for i := 1; i < len(values); i++ {

		ptr.Next = &ListElement{Value: values[i]}
		ptr.Next.Prev = ptr
		ptr = ptr.Next
	}

	return head
}
