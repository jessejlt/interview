package interview

import "sync"

// NewLFU creates an LFU cache of fixed size
// where all operations are O(1). LFU is threadsafe.
func NewLFU(size int) *LFU {

	n := &accessNode{value: 0}
	return &LFU{
		values: make(map[string]*valueNode, size),
		head:   n,
		cap:    size,
	}
}

// LFU is an LFU cache
type LFU struct {
	values map[string]*valueNode
	head   *accessNode
	cap    int
	len    int
	sync.RWMutex
}

// Set a value
func (l *LFU) Set(key string, value interface{}) {
	l.Lock()
	defer l.Unlock()

	// If we already have this key
	if v, ok := l.values[key]; ok {

		// Update the value
		v.value = value
		if v.frequency.next == nil {

			// Create a new node at the end of the list,
			// advance its counter, re-assign value node.
			v.frequency.next = &accessNode{
				value: v.frequency.value + 1,
				head:  v,
			}
			v.next = nil
			v.prev = nil
			v.frequency = v.frequency.next
		} else {

			// Advance node in frequency list
			v.frequency = v.frequency.next
			if v.frequency.head == nil {
				v.frequency.head = v
				v.next = nil
				v.prev = nil
			} else {
				v.next = v.frequency.head
				v.prev = nil
				v.frequency.head = v
			}
		}

		return
	}

	// If we available space
	if l.len < l.cap {
		l.len++

		// Ensure zero access node
		if l.head.value != 0 {

			n := &accessNode{
				value: 0,
				next:  l.head,
			}
			l.head = n

			// Assign value
			l.values[key] = &valueNode{
				value:     value,
				frequency: l.head,
			}

			return
		}

		l.values[key] = &valueNode{
			value:     value,
			frequency: l.head,
			next:      l.head.head,
		}
		l.head.head = l.values[key]

		return
	}

	// Evict

}

// Get the value for a key if it exists, else ok is false
func (l *LFU) Get(key string) (value interface{}, ok bool) {
	return
}

type accessNode struct {
	value int
	next  *accessNode
	head  *valueNode
}

type valueNode struct {
	key       string
	value     interface{}
	prev      *valueNode
	next      *valueNode
	frequency *accessNode
}
