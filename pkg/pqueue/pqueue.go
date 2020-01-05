package queue

import (
	"bytes"
	"container/heap"
	"fmt"
	"sync"
	"time"
)

type Item struct {
	Value    interface{} // The value of the item; arbitrary.
	priority int         // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index     int   // The index of the item in the heap.
	timestamp int64 //insertion timestamp
}

type priorityQueue []Item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq *priorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if (*pq)[i].priority == (*pq)[j].priority {
		return (*pq)[i].timestamp > (*pq)[j].timestamp && (*pq)[i].index > (*pq)[j].index
	}
	return (*pq)[i].priority > (*pq)[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	if pq.Len() > 0 {
		pq[i], pq[j] = pq[j], pq[i]
		pq[i].index = i
		pq[j].index = j
	}
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(Item)
	item.index = n
	if item.timestamp != -1 {
		item.timestamp = time.Now().UnixNano()
	}
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() (x interface{}) {
	l := pq.Len()
	if l > 0 {
		x = (*pq)[l-1]
		(*pq)[l-1] = Item{}
		*pq = (*pq)[:l-1]
	}
	return
}

func (pq *priorityQueue) Front() Item {
	old := *pq
	n := len(*pq)
	item := old[n-1]
	return item
}
func (pq *priorityQueue) Get(i int) Item {

	if i >= 0 && i < pq.Len() {
		return (*pq)[i]
	}
	return Item{}
}

// update modifies the priority and value of an Item in the queue.
func (pq *priorityQueue) update(item *Item, value string, priority int) {
	item.Value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// Listener is the function type to run on events.
type Listener func(interface{})

// ItemQueue the queue of Items
type Queue struct {
	items     *priorityQueue
	cond      sync.RWMutex
	front     int
	back      int
	length    int
	events    chan Item
	quit      chan bool
	listeners []Listener
}

// sendEvent send one or more events to the observer listeners.
func (q *Queue) sendEvent(event Item) {
	// NOTE: we do not lock this function directly.
	//
	// All functions using sendEvent must be locked
	// for operations using o.listeners.
	for _, listener := range q.listeners {
		go listener(event)
	}
}

// AddListener adds a listener function to run on event,
// the listener function will recive the event object as argument.
func (q *Queue) AddListener(l Listener) {
	// Check for mutex
	q.lazyInit()
	q.cond.Lock()
	defer q.cond.Unlock()
	q.listeners = append(q.listeners, l)
}

// handleEvent handle an event.
func (q *Queue) handleEvent(event Item, f *string) {
	// Lock:
	// 1. operations on listeners array (sendEvent).
	// 2. operations on bufferEvents array.
	// 3. operations using the watchPatterns set (matchFile).
	q.cond.Lock()
	q.sendEvent(event)
	defer q.cond.Unlock()
}

// eventLoop runs the event loop.
func (q *Queue) eventLoop() error {
	// Run observer.
	go func() {
		for {
			select {
			case event := <-q.events:
				q.handleEvent(event, nil)
			case <-q.quit:
				return
			}
		}
	}()

	return nil
}

// New returns an initialized empty queue.
func New() *Queue {
	return new(Queue).Init()
}

// Init initializes or clears queue q.
func (q *Queue) Init() *Queue {
	q.items = new(priorityQueue)
	q.front, q.back, q.length = 0, 0, 0
	heap.Init(q.items)
	if q.events != nil {
		return nil
	}
	q.events = make(chan Item)
	q.quit = make(chan bool)
	q.eventLoop()
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
	return q.items.Len()
}

// full returns true if the queue q is at capacity.
func (q *Queue) full() bool {
	return q.length == q.items.Len()
}

// empty returns true if the queue q has no elements.
func (q *Queue) empty() bool {
	return q.length == 0
}

// sparse returns true if the queue q has excess capacity.
func (q *Queue) sparse() bool {
	return 1 < q.length && q.length < q.items.Len()/4
}

// String returns a string representation of queue q formatted
// from front to back.
func (q *Queue) String() string {
	var result bytes.Buffer
	result.WriteByte('[')
	for i := 0; i < q.length; i++ {
		result.WriteString(fmt.Sprintf("%v", q.items.Get(i)))
		if i < q.length-1 {
			result.WriteByte(' ')
		}
	}
	result.WriteByte(']')
	return result.String()
}

// inc returns the next integer position wrapping around queue q.
func (q *Queue) inc(i int) int {
	return (i + 1) & (q.items.Len() - 1) // requires l = 2^n
}

// dec returns the previous integer position wrapping around queue q.
func (q *Queue) dec(i int) int {
	return (i - 1) & (q.items.Len() - 1) // requires l = 2^n
}

// Front returns the first element of queue q or nil.
func (q *Queue) Front() Item {
	// no need to check q.empty(), unused slots are nil
	defer q.cond.Unlock()
	q.cond.Lock()
	return q.items.Get(q.front)
}

// Back returns the last element of queue q or nil.
func (q *Queue) Back() interface{} {
	// no need to check q.empty(), unused slots are nil
	defer q.cond.Unlock()
	q.cond.Lock()
	return q.items.Get(q.dec(q.back))
}

// Enqueue inserts a new value v at the front of queue q.
func (q *Queue) Enqueue(item Item) {
	q.cond.Lock()
	q.lazyInit()
	// q.front = q.inc(q.front)
	heap.Push(q.items, item)
	q.cond.Unlock()
	q.events <- item
	q.length++
	return
}

// Dequeue removes and returns the first element of queue q or nil.
func (q *Queue) Dequeue() Item {
	q.cond.Lock()
	defer q.cond.Unlock()
	item := heap.Pop(q.items)
	if item != nil {
		q.front = q.dec(q.front)
		q.length--
		return item.(Item)
	}
	return Item{Value: nil}
}