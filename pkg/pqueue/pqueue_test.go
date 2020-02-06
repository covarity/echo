package queue

// References
// 1. https://golang.org/pkg/container/heap/#example__priorityQueue
// 2. https://golang.org/src/container/heap/heap.go
// 3. https://flaviocopes.com/golang-event-listeners/
// 4. https://github.com/nu7hatch/gopqueue/blob/master/pqueue.go
import (
	"math/rand"
	"testing"
	"time"
)

func ensureEmpty(t *testing.T, q *Queue) {
	if l := q.Len(); l != 0 {
		t.Errorf("q.Len() = %d, want %d", l, 0)
	}
	if e := q.Front(); e.Value != nil {
		t.Errorf("q.Front() = %v, want %v", e, nil)
	}
	if e := q.Dequeue(); e.Value != nil {
		t.Errorf("q.PopFront() = %v, want %v", e, nil)
	}
}

func TestNew(t *testing.T) {
	q := New(0)
	ensureEmpty(t, q)
}

func ensureSingleton(t *testing.T, q *Queue) {
	if l := q.Len(); l != 1 {
		t.Errorf("q.Len() = %d, want %d", l, 1)
	}
	if e := q.Front(); e.Value != 42 {
		t.Errorf("q.Front() = %v, want %v", e, 42)
	}
}

func TestSingleton(t *testing.T) {
	q := New(0)
	q.AddListener(func(e interface{}) {
		t.Logf("Received: %s.\n", e)
	})
	ensureEmpty(t, q)
	q.Enqueue(Item{Value: 42, priority: 0})
	ensureSingleton(t, q)
	q.Dequeue()
	ensureEmpty(t, q)
}

func TestDuos(t *testing.T) {
	q := New(0)
	q.AddListener(func(e interface{}) {
		println("Received: %s.\n", e)
	})
	ensureEmpty(t, q)
	q.Enqueue(Item{Value: 42, priority: 1})
	ensureSingleton(t, q)
	q.Enqueue(Item{Value: 43, priority: 0})
	if l := q.Len(); l != 2 {
		t.Errorf("q.Len() = %d, want %d", l, 2)
	}
	if e := q.Dequeue(); e.Value != 42 {
		t.Errorf("q.Front() = %v, want %v", e, 42)
	}
	if e := q.Dequeue(); e.Value != 43 {
		t.Errorf("q.Back() = %v, want %v", e, 43)
	}
}

func ensureLength(t *testing.T, q *Queue, len int) {
	if l := q.Len(); l != len {
		t.Errorf("q.Len() = %d, want %d", l, len)
	}
}

func TestAddListener(t *testing.T) {
	var output Item

	q := New(0)

	done := make(chan bool)
	defer close(done)

	q.AddListener(func(e interface{}) {
		output = e.(Item)
		done <- true
	})
	q.Enqueue(Item{Value: 43, priority: 0})

	<-done // blocks until listener is triggered

	if output.Value != 43 {
		t.Error("error recieving event")
	}
}

func TestZeroValue(t *testing.T) {
	var q Queue
	ts := int64(-1)
	q.Enqueue(Item{Value: 1, priority: 0, timestamp: ts})
	ensureLength(t, &q, 1)
	q.Enqueue(Item{Value: 2, priority: 0, timestamp: ts})
	ensureLength(t, &q, 2)
	q.Enqueue(Item{Value: 3, priority: 0, timestamp: ts})
	ensureLength(t, &q, 3)
	q.Enqueue(Item{Value: 4, priority: 0, timestamp: ts})
	ensureLength(t, &q, 4)
	q.Enqueue(Item{Value: 5, priority: 0, timestamp: ts})
	ensureLength(t, &q, 5)
	q.Enqueue(Item{Value: 6, priority: 0, timestamp: ts})
	ensureLength(t, &q, 6)
	q.Enqueue(Item{Value: 7, priority: 0, timestamp: ts})
	ensureLength(t, &q, 7)
	q.Enqueue(Item{Value: 8, priority: 0, timestamp: ts})
	ensureLength(t, &q, 8)
	q.Enqueue(Item{Value: 9, priority: 0, timestamp: ts})
	ensureLength(t, &q, 9)
	const want = "[{1 0 0 -1} {2 0 1 -1} {3 0 2 -1} {4 0 3 -1} {5 0 4 -1} {6 0 5 -1} {7 0 6 -1} {8 0 7 -1} {9 0 8 -1}]"
	if s := q.String(); s != want {
		t.Errorf("q.String() = %s, want %s", s, want)
	}
}

const size = 100

func TestGrowShrinkFILO(t *testing.T) {
	var q Queue
	for i := 0; i < size; i++ {
		q.Enqueue(Item{Value: i, priority: i})
		ensureLength(t, &q, i+1)
	}
	for i := size - 1; q.Len() > 0; i-- {
		x := q.Dequeue()
		if x.Value != i {
			t.Errorf("q.Dequeue() = %d, want %d", x, i)
		}
		ensureLength(t, &q, i)
	}
}
func TestGrowShrinkFIFO(t *testing.T) {
	var q Queue
	for i := 0; i < size; i++ {
		q.Enqueue(Item{Value: i, priority: size - i})
		ensureLength(t, &q, i+1)
	}
	for i := 0; q.Len() > 0; i++ {
		x := q.Dequeue()
		if x.Value != i {
			t.Errorf("q.Dequeue() = %d, want %d", x, i)
		}
		ensureLength(t, &q, size-i-1)
	}
}

func TestGrowShrinkSamePriority(t *testing.T) {
	var q Queue
	for i := 0; i < size; i++ {
		time.Sleep(1000)
		q.Enqueue(Item{Value: i, priority: 0})
		ensureLength(t, &q, i+1)
	}
	// println(q.String())
	for i := size - 1; q.Len() > 0; i-- {
		x := q.Dequeue()
		if x.Value != i {
			t.Errorf("q.Dequeue() = %d, want %d", x, i)
		}
		ensureLength(t, &q, i)
	}
}

// const size = 1024

func BenchmarkEnqueuePQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q Queue
		for n := 0; n < size; n++ {
			q.Enqueue(Item{Value: n, priority: 0})
		}
	}
}

var rands []float32

func makeRands() {
	if rands != nil {
		return
	}
	rand.Seed(64738)
	for i := 0; i < 4*size; i++ {
		rands = append(rands, rand.Float32())
	}
}
func BenchmarkRandomPQueue(b *testing.B) {
	makeRands()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var q Queue
		for n := 0; n < 4*size; n += 4 {
			if rands[n+1] < 0.8 {
				q.Enqueue(Item{Value: n, priority: 0})
			}
			if rands[n+2] < 0.5 {
				q.Dequeue()
			}
		}
	}
}

func BenchmarkGrowShrinkPQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q Queue
		for n := 0; n < size; n++ {
			q.Enqueue(Item{Value: i, priority: 0})
		}
		for n := 0; n < size; n++ {
			q.Dequeue()
		}
	}
}
