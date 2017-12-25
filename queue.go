package interview

// NewQueue creates a Queue
func NewQueue() *Queue {

	return new(Queue)
}

// Queue is an integer Queue
type Queue struct {
	head *queueNode
	tail *queueNode
}

type queueNode struct {
	data int
	next *queueNode
}

// Add an item to the end of the list
func (q *Queue) Add(data int) {

	if q == nil {
		q = NewQueue()
	}

	node := &queueNode{data: data}

	if q.head == nil {

		q.head = node
		q.tail = node
	} else {

		q.tail.next = node
		q.tail = node
	}
}

// Remove the first item in the list
func (q *Queue) Remove() int {

	if q == nil {
		q = NewQueue()
	}

	if q.head == nil {
		return -1
	}

	v := q.head.data
	q.head = q.head.next

	return v
}

// Peek returns the top of the queue
func (q *Queue) Peek() int {

	if q == nil {
		q = NewQueue()
	}

	if q.head == nil {
		return -1
	}

	return q.head.data
}

// IsEmpty returns ture if the queue is empty
func (q *Queue) IsEmpty() bool {

	if q == nil {
		q = NewQueue()
	}

	return q.head == nil
}
