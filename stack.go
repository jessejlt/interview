package interview

// NewStack creates a Stack
func NewStack(data []int) *Stack {

	s := new(Stack)
	for _, value := range data {

		s.Push(value)
	}

	return s
}

// Stack is an integer Stack
type Stack struct {
	top *stackNode
}

type stackNode struct {
	data int
	next *stackNode
}

// Pop removes the top element
func (s *Stack) Pop() int {

	if s == nil {
		s = NewStack(nil)
	}
	if s.top == nil {
		return -1
	}

	v := s.top.data
	s.top = s.top.next

	return v
}

// Push into the top element
func (s *Stack) Push(data int) {

	if s == nil {
		s = NewStack(nil)
	}

	node := &stackNode{data: data}

	if s.top != nil {

		node.next = s.top
	}

	s.top = node
}

// Peek returns the top element
func (s *Stack) Peek() int {

	if s == nil {
		s = NewStack(nil)
	}

	if s.top == nil {
		return -1
	}

	return s.top.data
}

// IsEmpty returns true if the Stack is empty
func (s *Stack) IsEmpty() bool {

	if s == nil {
		s = NewStack(nil)
	}

	return s.top == nil
}
