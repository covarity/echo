package queue

import (
	"bytes"
	"fmt"
	"sync"
)

type Item struct {
	value    interface{}
	priority int
}

// ItemQueue the queue of Items
type Queue struct {
	items  []Item
	lock   sync.RWMutex
	front  int
	back   int
	length int
}

// New returns an initialized empty queue.
func New() *Queue {
	return new(Queue).Init()
}

// Init initializes or clears queue q.
func (q *Queue) Init() *Queue {
	q.items = make([]Item, 1)
	q.front, q.back, q.length = 0, 0, 0
	return q
}

// lazyInit lazily initializes a zero Queue value.
//
// I am mostly doing this because container/list does the same thing.
// Personally I think it's a little wasteful because every single
// PushFront/PushBack is going to pay the overhead of calling this.
// But that's the price for making zero values useful immediately.
func (q *Queue) lazyInit() {
	if q.items == nil {
		q.Init()
	}
}

// Len returns the number of elements of queue q.
func (q *Queue) Len() int {
	return q.length
}

// full returns true if the queue q is at capacity.
func (q *Queue) full() bool {
	return q.length == len(q.items)
}

// empty returns true if the queue q has no elements.
func (q *Queue) empty() bool {
	return q.length == 0
}

// sparse returns true if the queue q has excess capacity.
func (q *Queue) sparse() bool {
	return 1 < q.length && q.length < len(q.items)/4
}

// resize adjusts the size of queue q's underlying slice.
func (q *Queue) resize(size int) {
	adjusted := make([]Item, size)
	if q.front < q.back {
		// items not "wrapped" around, one copy suffices
		copy(adjusted, q.items[q.front:q.back])
	} else {
		// items is "wrapped" around, need two copies
		n := copy(adjusted, q.items[q.front:])
		copy(adjusted[n:], q.items[:q.back])
	}
	q.items = adjusted
	q.front = 0
	q.back = q.length
}

// lazyGrow grows the underlying slice if necessary.
func (q *Queue) lazyGrow() {
	if q.full() {
		q.resize(len(q.items) * 2)
	}
}

// lazyShrink shrinks the underlying slice if advisable.
func (q *Queue) lazyShrink() {
	if q.sparse() {
		q.resize(len(q.items) / 2)
	}
}

// String returns a string representation of queue q formatted
// from front to back.
func (q *Queue) String() string {
	var result bytes.Buffer
	result.WriteByte('[')
	j := q.front
	for i := 0; i < q.length; i++ {
		result.WriteString(fmt.Sprintf("%v", q.items[j]))
		if i < q.length-1 {
			result.WriteByte(' ')
		}
		j = q.inc(j)
	}
	result.WriteByte(']')
	return result.String()
}

// inc returns the next integer position wrapping around queue q.
func (q *Queue) inc(i int) int {
	return (i + 1) & (len(q.items) - 1) // requires l = 2^n
}

// dec returns the previous integer position wrapping around queue q.
func (q *Queue) dec(i int) int {
	return (i - 1) & (len(q.items) - 1) // requires l = 2^n
}

// Front returns the first element of queue q or nil.
func (q *Queue) Front() Item {
	// no need to check q.empty(), unused slots are nil
	defer q.lock.RUnlock()
	q.lock.RLock()
	return q.items[q.front]
}

// Back returns the last element of queue q or nil.
func (q *Queue) Back() Item {
	// no need to check q.empty(), unused slots are nil
	defer q.lock.RUnlock()
	q.lock.RLock()
	return q.items[q.dec(q.back)]
}

// PushFront inserts a new value v at the front of queue q.
func (q *Queue) PushFront(item Item) {
	q.lock.Lock()
	q.lazyInit()
	q.lazyGrow()
	q.front = q.dec(q.front)
	q.items[q.front] = item
	q.length++
	q.lock.Unlock()
}

// PushBack inserts a new value v at the back of queue q.
func (q *Queue) PushBack(item Item) {
	q.lock.Lock()
	q.lazyInit()
	q.lazyGrow()
	q.items[q.back] = item
	q.back = q.inc(q.back)
	q.length++
	q.lock.Unlock()
}

// PopFront removes and returns the first element of queue q or nil.
func (q *Queue) PopFront() Item {
	q.lock.Lock()
	if q.empty() {
		q.lock.Unlock()
		return Item{value: nil, priority: 0}
	}
	v := q.items[q.front]
	q.items[q.front] = Item{value: nil, priority: 0} // unused slots must be nil
	q.front = q.inc(q.front)
	q.length--
	q.lazyShrink()
	q.lock.Unlock()
	return v
}

// PopBack removes and returns the last element of queue q or nil.
func (q *Queue) PopBack() Item {
	q.lock.Lock()
	if q.empty() {
		q.lock.Unlock()
		return Item{value: nil, priority: 0}
	}
	q.back = q.dec(q.back)
	v := q.items[q.back]
	q.items[q.back] = Item{value: nil, priority: 0} // unused slots must be nil
	q.length--
	q.lazyShrink()
	q.lock.Unlock()
	return v
}
