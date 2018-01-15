package interview

// NewStack creates a Stack
func NewStack() *Stack {

	return new(Stack)
}

// Stack is an integer Stack
type Stack struct {
	top *stackNode
}

type stackNode struct {
	data interface{}
	next *stackNode
}

// Pop removes the top element
func (s *Stack) Pop() interface{} {

	if s == nil {
		s = NewStack()
	}
	if s.top == nil {
		return -1
	}

	v := s.top.data
	s.top = s.top.next

	return v
}

// Push into the top element
func (s *Stack) Push(data interface{}) {

	if s == nil {
		s = NewStack()
	}

	node := &stackNode{data: data}

	if s.top != nil {

		node.next = s.top
	}

	s.top = node
}

// Peek returns the top element
func (s *Stack) Peek() interface{} {

	if s == nil {
		s = NewStack()
	}

	if s.top == nil {
		return -1
	}

	return s.top.data
}

// IsEmpty returns true if the Stack is empty
func (s *Stack) IsEmpty() bool {

	if s == nil {
		s = NewStack()
	}

	return s.top == nil
}
